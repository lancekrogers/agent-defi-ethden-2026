package agent

import (
	"fmt"
	"os"
	"time"

	"github.com/lancekrogers/agent-defi-ethden-2026/internal/base/trading"
	"github.com/lancekrogers/agent-defi-ethden-2026/internal/hcs"
)

// BaseConfig holds Base blockchain connection configuration.
type BaseConfig struct {
	// RPCURL is the Base Sepolia JSON-RPC endpoint.
	RPCURL string

	// ChainID is the target chain ID (84532 for Base Sepolia, 8453 for mainnet).
	ChainID int64

	// WalletAddress is this agent's Ethereum address.
	WalletAddress string

	// PrivateKey is the hex-encoded private key for signing transactions.
	PrivateKey string

	// ERC8004ContractAddress is the ERC-8004 identity registry contract.
	ERC8004ContractAddress string

	// BuilderCode is the ERC-8021 attribution code (Ethereum address, 20 bytes).
	BuilderCode string

	// DEXRouterAddress is the Uniswap v3 (or compatible) router address.
	DEXRouterAddress string
}

// HCSConfig holds Hedera Consensus Service configuration.
type HCSConfig struct {
	// TaskTopicID is the HCS topic for receiving task assignments.
	TaskTopicID string

	// ResultTopicID is the HCS topic for publishing results and reports.
	ResultTopicID string
}

// TradingConfig holds trading strategy and execution configuration.
type TradingConfig struct {
	// TokenIn is the address of the token to sell (e.g., USDC on Base Sepolia).
	TokenIn string

	// TokenOut is the address of the token to buy (e.g., WETH on Base Sepolia).
	TokenOut string

	// BuyThreshold is the fractional price deviation that triggers a buy signal.
	BuyThreshold float64

	// SellThreshold is the fractional price deviation that triggers a sell signal.
	SellThreshold float64

	// MaxPositionSize is the maximum single trade size in base asset units.
	MaxPositionSize float64

	// MinLiquidity is the minimum required DEX pool liquidity in USD.
	MinLiquidity float64
}

// Config holds all configuration for the DeFi agent.
type Config struct {
	// AgentID is the unique identifier for this agent instance.
	AgentID string

	// DaemonAddr is the address of the daemon client for coordination.
	DaemonAddr string

	// HealthInterval is how often to send heartbeat messages.
	HealthInterval time.Duration

	// TradingInterval is how often to evaluate the strategy and execute trades.
	TradingInterval time.Duration

	// Base holds Base chain connection configuration.
	Base BaseConfig

	// HCS holds Hedera Consensus Service configuration.
	HCS HCSConfig

	// Trading holds trading strategy configuration.
	Trading TradingConfig
}

// HCSHandler builds an HCS handler config from the agent config.
func (c *Config) HCSHandler(transport hcs.Transport) hcs.HandlerConfig {
	return hcs.HandlerConfig{
		Transport:     transport,
		TaskTopicID:   c.HCS.TaskTopicID,
		ResultTopicID: c.HCS.ResultTopicID,
		AgentID:       c.AgentID,
	}
}

// StrategyConfig builds a MeanReversionConfig from the agent's trading config.
func (c *Config) StrategyConfig() trading.MeanReversionConfig {
	return trading.MeanReversionConfig{
		TokenIn:         c.Trading.TokenIn,
		TokenOut:        c.Trading.TokenOut,
		BuyThreshold:    c.Trading.BuyThreshold,
		SellThreshold:   c.Trading.SellThreshold,
		MaxPositionSize: c.Trading.MaxPositionSize,
		MinLiquidity:    c.Trading.MinLiquidity,
	}
}

// LoadConfig reads DeFi agent configuration from environment variables.
// All env vars use the DEFI_ prefix.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	cfg.AgentID = os.Getenv("DEFI_AGENT_ID")
	if cfg.AgentID == "" {
		return nil, fmt.Errorf("config: DEFI_AGENT_ID is required")
	}

	cfg.DaemonAddr = envOr("DEFI_DAEMON_ADDR", "localhost:9090")

	healthStr := os.Getenv("DEFI_HEALTH_INTERVAL")
	if healthStr == "" {
		cfg.HealthInterval = 30 * time.Second
	} else {
		dur, err := time.ParseDuration(healthStr)
		if err != nil {
			return nil, fmt.Errorf("config: invalid DEFI_HEALTH_INTERVAL: %w", err)
		}
		cfg.HealthInterval = dur
	}

	tradingStr := os.Getenv("DEFI_TRADING_INTERVAL")
	if tradingStr == "" {
		cfg.TradingInterval = 60 * time.Second
	} else {
		dur, err := time.ParseDuration(tradingStr)
		if err != nil {
			return nil, fmt.Errorf("config: invalid DEFI_TRADING_INTERVAL: %w", err)
		}
		cfg.TradingInterval = dur
	}

	// Base chain configuration.
	cfg.Base.RPCURL = envOr("DEFI_BASE_RPC_URL", "https://sepolia.base.org")
	cfg.Base.ChainID = 84532 // Base Sepolia default
	cfg.Base.WalletAddress = os.Getenv("DEFI_WALLET_ADDRESS")
	cfg.Base.PrivateKey = os.Getenv("DEFI_PRIVATE_KEY")
	cfg.Base.ERC8004ContractAddress = os.Getenv("DEFI_ERC8004_CONTRACT")
	cfg.Base.BuilderCode = os.Getenv("DEFI_BUILDER_CODE")
	cfg.Base.DEXRouterAddress = envOr("DEFI_DEX_ROUTER", "0x0000000000000000000000000000000000000000")

	// HCS configuration.
	cfg.HCS.TaskTopicID = os.Getenv("HCS_TASK_TOPIC")
	cfg.HCS.ResultTopicID = os.Getenv("HCS_RESULT_TOPIC")

	// Trading strategy configuration.
	cfg.Trading.TokenIn = envOr("DEFI_TOKEN_IN", "0x036CbD53842c5426634e7929541eC2318f3dCF7e")  // USDC on Base Sepolia
	cfg.Trading.TokenOut = envOr("DEFI_TOKEN_OUT", "0x4200000000000000000000000000000000000006") // WETH on Base Sepolia
	cfg.Trading.BuyThreshold = 0.02
	cfg.Trading.SellThreshold = 0.02
	cfg.Trading.MaxPositionSize = 0.001 // 0.001 WETH max per trade
	cfg.Trading.MinLiquidity = 10_000   // $10k minimum liquidity

	return cfg, nil
}

func envOr(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
