package trading

import (
	"context"
	"fmt"
	"time"
)

// Strategy defines the interface for a DeFi trading strategy.
// Implementations analyze market state and produce trading signals.
type Strategy interface {
	// Name returns the human-readable name of this strategy.
	Name() string

	// Evaluate analyzes current market conditions and returns a trading signal.
	// Returns ErrMarketDataUnavailable if the market state is stale or invalid.
	Evaluate(ctx context.Context, market MarketState) (*Signal, error)

	// MaxPosition returns the maximum position size in base asset units.
	// No single trade should exceed this size.
	MaxPosition() float64
}

// MeanReversionConfig holds parameters for the mean reversion strategy.
type MeanReversionConfig struct {
	// TokenIn is the address of the token to sell (e.g., USDC).
	TokenIn string

	// TokenOut is the address of the token to buy (e.g., WETH).
	TokenOut string

	// BuyThreshold is the percentage below the moving average that triggers a buy.
	// Example: 0.02 means buy when price is 2% below moving average.
	BuyThreshold float64

	// SellThreshold is the percentage above the moving average that triggers a sell.
	// Example: 0.02 means sell when price is 2% above moving average.
	SellThreshold float64

	// MaxPositionSize is the maximum trade size in base asset units.
	MaxPositionSize float64

	// MinLiquidity is the minimum required DEX liquidity in USD to trade.
	MinLiquidity float64

	// DataStalenessLimit is the maximum age of market data before it is
	// considered too stale to trade on.
	DataStalenessLimit time.Duration
}

// MeanReversionStrategy produces buy signals when price is below the moving
// average by more than BuyThreshold, and sell signals when price is above
// the moving average by more than SellThreshold.
//
// Mean reversion is based on the observation that asset prices tend to
// revert to their historical mean over time.
type MeanReversionStrategy struct {
	cfg MeanReversionConfig
}

// NewMeanReversionStrategy creates a mean reversion strategy with the given config.
func NewMeanReversionStrategy(cfg MeanReversionConfig) *MeanReversionStrategy {
	if cfg.BuyThreshold == 0 {
		cfg.BuyThreshold = 0.02 // 2% below MA = buy
	}
	if cfg.SellThreshold == 0 {
		cfg.SellThreshold = 0.02 // 2% above MA = sell
	}
	if cfg.DataStalenessLimit == 0 {
		cfg.DataStalenessLimit = 5 * time.Minute
	}
	return &MeanReversionStrategy{cfg: cfg}
}

// Name returns the strategy identifier.
func (s *MeanReversionStrategy) Name() string {
	return "mean_reversion"
}

// MaxPosition returns the configured maximum position size.
func (s *MeanReversionStrategy) MaxPosition() float64 {
	return s.cfg.MaxPositionSize
}

// Evaluate analyzes the market state against the moving average and produces
// a buy, sell, or hold signal based on the configured thresholds.
func (s *MeanReversionStrategy) Evaluate(ctx context.Context, market MarketState) (*Signal, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("strategy: context cancelled before evaluate: %w", err)
	}

	if err := s.validateMarket(market); err != nil {
		return nil, err
	}

	signal := &Signal{
		GeneratedAt: time.Now(),
		TokenIn:     s.cfg.TokenIn,
		TokenOut:    s.cfg.TokenOut,
	}

	deviation := s.priceDeviation(market.Price, market.MovingAverage)

	switch {
	case deviation < -s.cfg.BuyThreshold:
		// Price is below the moving average by more than the buy threshold.
		signal.Type = SignalBuy
		signal.Confidence = s.confidence(deviation, s.cfg.BuyThreshold)
		signal.SuggestedSize = s.positionSize(signal.Confidence)
		signal.Reason = fmt.Sprintf("price %.4f is %.2f%% below MA %.4f",
			market.Price, -deviation*100, market.MovingAverage)

	case deviation > s.cfg.SellThreshold:
		// Price is above the moving average by more than the sell threshold.
		signal.Type = SignalSell
		signal.Confidence = s.confidence(deviation, s.cfg.SellThreshold)
		signal.SuggestedSize = s.positionSize(signal.Confidence)
		signal.Reason = fmt.Sprintf("price %.4f is %.2f%% above MA %.4f",
			market.Price, deviation*100, market.MovingAverage)

	default:
		signal.Type = SignalHold
		signal.Confidence = 1.0 - abs(deviation)/s.cfg.BuyThreshold
		signal.Reason = fmt.Sprintf("price %.4f within %.2f%% of MA %.4f",
			market.Price, abs(deviation)*100, market.MovingAverage)
	}

	return signal, nil
}

// validateMarket checks market data for validity and staleness.
func (s *MeanReversionStrategy) validateMarket(market MarketState) error {
	if market.Price <= 0 || market.MovingAverage <= 0 {
		return fmt.Errorf("strategy: %w: price or MA is zero", ErrMarketDataUnavailable)
	}

	if market.Liquidity < s.cfg.MinLiquidity && s.cfg.MinLiquidity > 0 {
		return fmt.Errorf("strategy: %w: liquidity %.2f below minimum %.2f",
			ErrInsufficientLiquidity, market.Liquidity, s.cfg.MinLiquidity)
	}

	age := time.Since(market.FetchedAt)
	if age > s.cfg.DataStalenessLimit {
		return fmt.Errorf("strategy: %w: market data is %v old (limit %v)",
			ErrMarketDataUnavailable, age, s.cfg.DataStalenessLimit)
	}

	return nil
}

// priceDeviation returns the fractional deviation of price from the moving average.
// Positive values mean price is above MA; negative means below.
func (s *MeanReversionStrategy) priceDeviation(price, ma float64) float64 {
	if ma == 0 {
		return 0
	}
	return (price - ma) / ma
}

// confidence converts a deviation into a confidence score capped at 1.0.
// Larger deviations from the threshold produce higher confidence.
func (s *MeanReversionStrategy) confidence(deviation, threshold float64) float64 {
	// Scale: at the threshold, confidence is 0.5; doubles confidence per threshold.
	c := abs(deviation) / threshold * 0.5
	if c > 1.0 {
		return 1.0
	}
	return c
}

// positionSize returns a trade size proportional to confidence, capped at MaxPosition.
func (s *MeanReversionStrategy) positionSize(confidence float64) float64 {
	return s.cfg.MaxPositionSize * confidence
}

// abs returns the absolute value of a float64.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
