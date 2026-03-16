package vault

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lancekrogers/agent-defi/internal/base/ethutil"
)

// Client reads vault state and submits swaps via executeSwap.
type Client interface {
	USDCBalance(ctx context.Context) (*big.Int, error)
	TotalAssets(ctx context.Context) (*big.Int, error)
	SharePrice(ctx context.Context) (*big.Float, error)
	ExecuteSwap(ctx context.Context, params SwapParams) (common.Hash, error)
	HeldTokens(ctx context.Context) ([]common.Address, error)
}

// SwapParams holds the parameters for a vault swap.
type SwapParams struct {
	TokenIn      common.Address
	TokenOut     common.Address
	AmountIn     *big.Int
	MinAmountOut *big.Int
	Reason       []byte
}

// Config holds vault client configuration.
type Config struct {
	RPCURL       string
	ChainID      int64
	VaultAddress string
	PrivateKey   string
}

type client struct {
	cfg Config
}

// NewClient creates a vault client.
func NewClient(cfg Config) Client {
	return &client{cfg: cfg}
}

func (c *client) USDCBalance(ctx context.Context) (*big.Int, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("vault: context cancelled: %w", err)
	}
	return big.NewInt(0), nil // placeholder — calls asset().balanceOf(vault) via eth_call
}

func (c *client) TotalAssets(ctx context.Context) (*big.Int, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("vault: context cancelled: %w", err)
	}
	return big.NewInt(0), nil // placeholder — calls vault.totalAssets() via eth_call
}

func (c *client) SharePrice(ctx context.Context) (*big.Float, error) {
	return new(big.Float), nil // placeholder — totalAssets / totalSupply
}

func (c *client) ExecuteSwap(ctx context.Context, params SwapParams) (common.Hash, error) {
	if err := ctx.Err(); err != nil {
		return common.Hash{}, fmt.Errorf("vault: context cancelled: %w", err)
	}

	key, err := ethutil.LoadKey(c.cfg.PrivateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("vault: load key: %w", err)
	}

	ethClient, err := ethutil.DialClient(ctx, c.cfg.RPCURL)
	if err != nil {
		return common.Hash{}, fmt.Errorf("vault: dial: %w", err)
	}
	defer ethClient.Close()

	vaultAddr := common.HexToAddress(c.cfg.VaultAddress)
	bound, err := NewObeyVault(vaultAddr, ethClient)
	if err != nil {
		return common.Hash{}, fmt.Errorf("vault: bind contract: %w", err)
	}

	opts, err := ethutil.MakeTransactOpts(ctx, key, c.cfg.ChainID)
	if err != nil {
		return common.Hash{}, fmt.Errorf("vault: make tx opts: %w", err)
	}

	tx, err := bound.ExecuteSwap(opts, params.TokenIn, params.TokenOut, params.AmountIn, params.MinAmountOut, params.Reason)
	if err != nil {
		return common.Hash{}, fmt.Errorf("vault: executeSwap failed: %w", err)
	}

	return tx.Hash(), nil
}

func (c *client) HeldTokens(ctx context.Context) ([]common.Address, error) {
	return nil, nil // placeholder — calls heldTokenCount + heldTokenAt via abigen
}
