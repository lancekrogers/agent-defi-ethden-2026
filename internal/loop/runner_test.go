package loop

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lancekrogers/agent-defi/internal/base/trading"
	"github.com/lancekrogers/agent-defi/internal/risk"
	"github.com/lancekrogers/agent-defi/internal/vault"
)

type mockVault struct {
	swapCalled bool
	lastParams vault.SwapParams
}

func (m *mockVault) USDCBalance(ctx context.Context) (*big.Int, error) {
	return big.NewInt(100_000_000), nil
}
func (m *mockVault) TotalAssets(ctx context.Context) (*big.Int, error) {
	return big.NewInt(100_000_000), nil
}
func (m *mockVault) SharePrice(ctx context.Context) (*big.Float, error) {
	return new(big.Float).SetFloat64(1.0), nil
}
func (m *mockVault) ExecuteSwap(ctx context.Context, params vault.SwapParams) (common.Hash, error) {
	m.swapCalled = true
	m.lastParams = params
	return common.HexToHash("0xabc123"), nil
}
func (m *mockVault) HeldTokens(ctx context.Context) ([]common.Address, error) {
	return nil, nil
}

type mockExecutor struct{}

func (m *mockExecutor) Execute(ctx context.Context, trade trading.Trade) (*trading.TradeResult, error) {
	return nil, nil
}
func (m *mockExecutor) GetBalance(ctx context.Context, tokenAddress string) (*trading.Balance, error) {
	return &trading.Balance{AmountWei: "0x0"}, nil
}
func (m *mockExecutor) GetMarketState(ctx context.Context, tokenIn, tokenOut string) (*trading.MarketState, error) {
	return &trading.MarketState{
		TokenIn:       tokenIn,
		TokenOut:      tokenOut,
		Price:         2500.0,
		MovingAverage: 2600.0,
		Liquidity:     1_000_000,
		FetchedAt:     time.Now(),
	}, nil
}

type mockStrategy struct {
	signal *trading.Signal
}

func (m *mockStrategy) Name() string        { return "mock" }
func (m *mockStrategy) MaxPosition() float64 { return 1000 }
func (m *mockStrategy) Evaluate(ctx context.Context, market trading.MarketState) (*trading.Signal, error) {
	return m.signal, nil
}

func TestRunner_BuySignalExecutesSwap(t *testing.T) {
	mv := &mockVault{}
	r := New(
		Config{
			Interval: time.Second,
			TokenIn:  common.HexToAddress("0x1111"),
			TokenOut: common.HexToAddress("0x2222"),
		},
		nil,
		mv,
		&mockExecutor{},
		&mockStrategy{signal: &trading.Signal{
			Type:          trading.SignalBuy,
			Confidence:    0.8,
			SuggestedSize: 50.0,
			Reason:        "test buy",
			TokenIn:       "0x1111",
			TokenOut:      "0x2222",
		}},
		risk.NewManager(risk.Config{MaxPositionUSD: 200000, MaxDailyVolumeUSD: 1000000}),
	)

	err := r.cycle(context.Background())
	if err != nil {
		t.Fatalf("cycle failed: %v", err)
	}
	if !mv.swapCalled {
		t.Fatal("expected vault.ExecuteSwap to be called on buy signal")
	}
}

func TestRunner_HoldSignalSkipsSwap(t *testing.T) {
	mv := &mockVault{}
	r := New(
		Config{Interval: time.Second},
		nil,
		mv,
		&mockExecutor{},
		&mockStrategy{signal: &trading.Signal{
			Type:   trading.SignalHold,
			Reason: "no signal",
		}},
		risk.NewManager(risk.Config{MaxPositionUSD: 200000, MaxDailyVolumeUSD: 1000000}),
	)

	err := r.cycle(context.Background())
	if err != nil {
		t.Fatalf("cycle failed: %v", err)
	}
	if mv.swapCalled {
		t.Fatal("vault.ExecuteSwap should NOT be called on hold signal")
	}
}
