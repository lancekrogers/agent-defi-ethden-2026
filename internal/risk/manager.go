package risk

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lancekrogers/agent-defi/internal/base/trading"
)

// Config holds risk management parameters.
type Config struct {
	MaxPositionUSD    float64
	MaxDailyVolumeUSD float64
	MaxDrawdownPct    float64
	InitialNAV        float64
}

// Manager enforces position sizing, daily volume caps, and drawdown halts.
type Manager struct {
	cfg         Config
	mu          sync.Mutex
	dailyVolume float64
	currentDay  int
	currentNAV  float64
	peakNAV     float64
}

// NewManager creates a risk manager with the given configuration.
func NewManager(cfg Config) *Manager {
	nav := cfg.InitialNAV
	if nav == 0 {
		nav = 1.0
	}
	return &Manager{
		cfg:        cfg,
		currentDay: time.Now().YearDay(),
		currentNAV: nav,
		peakNAV:    nav,
	}
}

// Check validates a proposed trade against risk limits.
func (m *Manager) Check(ctx context.Context, signal *trading.Signal, price float64) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("risk: context cancelled: %w", err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.resetDayIfNeeded()

	positionUSD := signal.SuggestedSize * price
	if positionUSD > m.cfg.MaxPositionUSD {
		return fmt.Errorf("risk: position %.2f USD exceeds max %.2f USD",
			positionUSD, m.cfg.MaxPositionUSD)
	}

	if m.dailyVolume+positionUSD > m.cfg.MaxDailyVolumeUSD {
		return fmt.Errorf("risk: daily volume %.2f + %.2f exceeds cap %.2f USD",
			m.dailyVolume, positionUSD, m.cfg.MaxDailyVolumeUSD)
	}

	if m.cfg.MaxDrawdownPct > 0 && m.peakNAV > 0 {
		drawdown := (m.peakNAV - m.currentNAV) / m.peakNAV
		if drawdown > m.cfg.MaxDrawdownPct {
			return fmt.Errorf("risk: drawdown %.2f%% exceeds halt threshold %.2f%%",
				drawdown*100, m.cfg.MaxDrawdownPct*100)
		}
	}

	return nil
}

// Clamp reduces the signal's suggested size to fit within position limits.
func (m *Manager) Clamp(signal *trading.Signal, price float64) {
	maxSize := m.cfg.MaxPositionUSD / price
	if signal.SuggestedSize > maxSize {
		signal.SuggestedSize = maxSize
	}
}

// RecordTrade tracks a completed trade against daily volume limits.
func (m *Manager) RecordTrade(size, price float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.resetDayIfNeeded()
	m.dailyVolume += size * price
}

// UpdateNAV updates the current net asset value for drawdown tracking.
func (m *Manager) UpdateNAV(nav float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentNAV = nav
	if nav > m.peakNAV {
		m.peakNAV = nav
	}
}

func (m *Manager) resetDayIfNeeded() {
	today := time.Now().YearDay()
	if today != m.currentDay {
		m.currentDay = today
		m.dailyVolume = 0
	}
}
