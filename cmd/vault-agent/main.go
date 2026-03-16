package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lancekrogers/agent-defi/internal/base/trading"
	"github.com/lancekrogers/agent-defi/internal/loop"
	"github.com/lancekrogers/agent-defi/internal/risk"
	"github.com/lancekrogers/agent-defi/internal/strategy"
	"github.com/lancekrogers/agent-defi/internal/vault"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	vaultCfg := vault.Config{
		RPCURL:       envOr("VAULT_RPC_URL", "https://sepolia.base.org"),
		ChainID:      84532,
		VaultAddress: os.Getenv("VAULT_ADDRESS"),
		PrivateKey:   os.Getenv("AGENT_PRIVATE_KEY"),
	}

	vaultClient := vault.NewClient(vaultCfg)

	executor := trading.NewExecutor(trading.ExecutorConfig{
		RPCURL:           vaultCfg.RPCURL,
		ChainID:          vaultCfg.ChainID,
		WalletAddress:    os.Getenv("AGENT_WALLET_ADDRESS"),
		PrivateKey:       vaultCfg.PrivateKey,
		DEXRouterAddress: envOr("DEX_ROUTER", "0x2626664c2603336E57B271c5C0b26F421741e481"),
	})

	llmClient := &strategy.ObeyClient{
		Socket:   envOr("OBEY_SOCKET", "/tmp/obey.sock"),
		Campaign: envOr("OBEY_CAMPAIGN", "Obey-Agent-Economy"),
		Provider: envOr("OBEY_PROVIDER", "claude-code"),
		Model:    envOr("LLM_MODEL", "claude-sonnet-4-6"),
		Festival: envOr("OBEY_FESTIVAL", "OV0001"),
	}

	strat := strategy.NewLLMStrategy(strategy.LLMConfig{
		TokenIn:         os.Getenv("TOKEN_IN"),
		TokenOut:        os.Getenv("TOKEN_OUT"),
		MaxPositionSize: 100.0,
	}, llmClient)

	riskMgr := risk.NewManager(risk.Config{
		MaxPositionUSD:    1000,
		MaxDailyVolumeUSD: 10000,
		MaxDrawdownPct:    0.10,
		InitialNAV:        10000,
	})

	runner := loop.New(loop.Config{
		Interval: 5 * time.Minute,
		TokenIn:  common.HexToAddress(os.Getenv("TOKEN_IN")),
		TokenOut: common.HexToAddress(os.Getenv("TOKEN_OUT")),
	}, log, vaultClient, executor, strat, riskMgr)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log.Info("vault agent starting",
		"vault", vaultCfg.VaultAddress,
		"strategy", strat.Name(),
	)

	if err := runner.Run(ctx); err != nil && err != context.Canceled {
		log.Error("agent exited with error", "error", err)
		os.Exit(1)
	}
}

func envOr(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
