package risk

import (
	"context"
	"testing"

	"github.com/lancekrogers/agent-defi/internal/base/trading"
)

func TestManager_RejectsOversizedPosition(t *testing.T) {
	m := NewManager(Config{MaxPositionUSD: 500, MaxDailyVolumeUSD: 10000})
	signal := &trading.Signal{
		Type:          trading.SignalBuy,
		SuggestedSize: 1000.0,
	}
	err := m.Check(context.Background(), signal, 2500.0)
	if err == nil {
		t.Fatal("expected rejection for oversized position")
	}
}

func TestManager_ClampsPosition(t *testing.T) {
	m := NewManager(Config{MaxPositionUSD: 500, MaxDailyVolumeUSD: 10000})
	signal := &trading.Signal{
		Type:          trading.SignalBuy,
		SuggestedSize: 1000.0,
	}
	m.Clamp(signal, 2500.0)
	if signal.SuggestedSize > 500.0/2500.0 {
		t.Errorf("expected clamped size <= %f, got %f", 500.0/2500.0, signal.SuggestedSize)
	}
}

func TestManager_DailyVolumeTracking(t *testing.T) {
	m := NewManager(Config{MaxPositionUSD: 10000, MaxDailyVolumeUSD: 1000})
	signal := &trading.Signal{Type: trading.SignalBuy, SuggestedSize: 0.5}

	err := m.Check(context.Background(), signal, 2000.0)
	if err != nil {
		t.Fatalf("first trade should pass: %v", err)
	}
	m.RecordTrade(0.5, 2000.0)

	err = m.Check(context.Background(), signal, 2000.0)
	if err == nil {
		t.Fatal("expected daily volume rejection")
	}
}

func TestManager_DrawdownHalt(t *testing.T) {
	m := NewManager(Config{
		MaxPositionUSD:    10000,
		MaxDailyVolumeUSD: 100000,
		MaxDrawdownPct:    0.10,
		InitialNAV:        10000,
	})
	m.UpdateNAV(8500)
	signal := &trading.Signal{Type: trading.SignalBuy, SuggestedSize: 0.1}

	err := m.Check(context.Background(), signal, 2500.0)
	if err == nil {
		t.Fatal("expected drawdown halt")
	}
}
