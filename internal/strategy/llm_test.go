package strategy

import (
	"context"
	"testing"
	"time"

	"github.com/lancekrogers/agent-defi/internal/base/trading"
)

type mockLLM struct {
	response string
	err      error
}

func (m *mockLLM) Complete(ctx context.Context, prompt string) (string, error) {
	return m.response, m.err
}

func TestLLMStrategy_BuySignal(t *testing.T) {
	s := NewLLMStrategy(LLMConfig{
		TokenIn:         "0xUSDC",
		TokenOut:        "0xWETH",
		MaxPositionSize: 100.0,
	}, &mockLLM{response: `{"action":"buy","confidence":0.8,"size":0.5,"reason":"ETH undervalued"}`})

	market := trading.MarketState{
		TokenIn:       "0xUSDC",
		TokenOut:      "0xWETH",
		Price:         2500.0,
		MovingAverage: 2600.0,
		Liquidity:     1000000,
		FetchedAt:     time.Now(),
	}

	signal, err := s.Evaluate(context.Background(), market)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if signal.Type != trading.SignalBuy {
		t.Errorf("expected buy signal, got %s", signal.Type)
	}
	if signal.Confidence < 0.5 {
		t.Errorf("expected confidence >= 0.5, got %f", signal.Confidence)
	}
}

func TestLLMStrategy_HoldSignal(t *testing.T) {
	s := NewLLMStrategy(LLMConfig{
		TokenIn:         "0xUSDC",
		TokenOut:        "0xWETH",
		MaxPositionSize: 100.0,
	}, &mockLLM{response: `{"action":"hold","confidence":0.3,"size":0,"reason":"no clear signal"}`})

	market := trading.MarketState{
		TokenIn:   "0xUSDC",
		TokenOut:  "0xWETH",
		Price:     2500.0,
		FetchedAt: time.Now(),
	}

	signal, err := s.Evaluate(context.Background(), market)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if signal.Type != trading.SignalHold {
		t.Errorf("expected hold signal, got %s", signal.Type)
	}
}

func TestLLMStrategy_ContextCancellation(t *testing.T) {
	s := NewLLMStrategy(LLMConfig{}, &mockLLM{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := s.Evaluate(ctx, trading.MarketState{})
	if err == nil {
		t.Fatal("expected error on cancelled context")
	}
}
