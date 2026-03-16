package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lancekrogers/agent-defi/internal/base/trading"
)

// LLMClient abstracts an LLM completion API for strategy evaluation.
type LLMClient interface {
	Complete(ctx context.Context, prompt string) (string, error)
}

// LLMConfig holds parameters for the LLM-driven trading strategy.
type LLMConfig struct {
	TokenIn         string
	TokenOut        string
	MaxPositionSize float64
}

// LLMStrategy uses an LLM to evaluate market conditions and produce signals.
type LLMStrategy struct {
	cfg LLMConfig
	llm LLMClient
}

// NewLLMStrategy creates a new LLM-driven trading strategy.
func NewLLMStrategy(cfg LLMConfig, llm LLMClient) *LLMStrategy {
	return &LLMStrategy{cfg: cfg, llm: llm}
}

// Name returns the strategy identifier.
func (s *LLMStrategy) Name() string { return "llm_momentum" }

// MaxPosition returns the configured maximum position size.
func (s *LLMStrategy) MaxPosition() float64 { return s.cfg.MaxPositionSize }

type llmResponse struct {
	Action     string  `json:"action"`
	Confidence float64 `json:"confidence"`
	Size       float64 `json:"size"`
	Reason     string  `json:"reason"`
}

// Evaluate sends market data to the LLM and parses the trading recommendation.
func (s *LLMStrategy) Evaluate(ctx context.Context, market trading.MarketState) (*trading.Signal, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("llm strategy: context cancelled: %w", err)
	}

	prompt := fmt.Sprintf(
		`You are a DeFi trading agent. Analyze this market data and decide whether to buy, sell, or hold.

Token pair: %s → %s
Current price: %.6f
Moving average: %.6f
Liquidity: %.2f

Respond with JSON only: {"action":"buy|sell|hold","confidence":0.0-1.0,"size":0.0-1.0,"reason":"..."}`,
		market.TokenIn, market.TokenOut,
		market.Price, market.MovingAverage, market.Liquidity,
	)

	resp, err := s.llm.Complete(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("llm strategy: completion failed: %w", err)
	}

	var parsed llmResponse
	if err := json.Unmarshal([]byte(resp), &parsed); err != nil {
		return nil, fmt.Errorf("llm strategy: parse response: %w", err)
	}

	signal := &trading.Signal{
		GeneratedAt: time.Now(),
		TokenIn:     s.cfg.TokenIn,
		TokenOut:    s.cfg.TokenOut,
		Confidence:  parsed.Confidence,
		Reason:      parsed.Reason,
	}

	switch parsed.Action {
	case "buy":
		signal.Type = trading.SignalBuy
		signal.SuggestedSize = parsed.Size * s.cfg.MaxPositionSize
	case "sell":
		signal.Type = trading.SignalSell
		signal.SuggestedSize = parsed.Size * s.cfg.MaxPositionSize
	default:
		signal.Type = trading.SignalHold
	}

	return signal, nil
}
