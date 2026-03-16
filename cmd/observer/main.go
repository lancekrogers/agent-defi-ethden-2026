package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/lancekrogers/agent-defi/internal/vault"
)

func main() {
	vaultAddr := os.Getenv("VAULT_ADDRESS")
	rpcURL := os.Getenv("RPC_URL")
	if vaultAddr == "" || rpcURL == "" {
		log.Fatal("VAULT_ADDRESS and RPC_URL required")
	}

	client := vault.NewClient(vault.Config{
		RPCURL:       rpcURL,
		VaultAddress: vaultAddr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("=== OBEY Vault Status ===")
	fmt.Printf("Vault: %s\n\n", vaultAddr)

	balance, err := client.USDCBalance(ctx)
	if err != nil {
		fmt.Printf("USDC Balance: error (%v)\n", err)
	} else {
		fmt.Printf("USDC Balance: %s\n", formatUSDC(balance))
	}

	total, err := client.TotalAssets(ctx)
	if err != nil {
		fmt.Printf("Total Assets (NAV): error (%v)\n", err)
	} else {
		fmt.Printf("Total Assets (NAV): %s USDC\n", formatUSDC(total))
	}

	sharePrice, err := client.SharePrice(ctx)
	if err != nil {
		fmt.Printf("Share Price: error (%v)\n", err)
	} else {
		fmt.Printf("Share Price: %s\n", sharePrice.Text('f', 6))
	}

	tokens, err := client.HeldTokens(ctx)
	if err == nil && len(tokens) > 0 {
		fmt.Println("\nHeld Tokens:")
		for _, t := range tokens {
			fmt.Printf("  - %s\n", t.Hex())
		}
	}

	fmt.Println("\n(Trade history via SwapExecuted events — coming soon)")
}

func formatUSDC(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	f := new(big.Float).SetInt(wei)
	f.Quo(f, new(big.Float).SetFloat64(1e6))
	return f.Text('f', 2)
}
