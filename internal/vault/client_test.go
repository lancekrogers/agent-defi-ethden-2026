package vault

import (
	"context"
	"testing"
)

func TestClientContextCancellation(t *testing.T) {
	c := NewClient(Config{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := c.USDCBalance(ctx)
	if err == nil {
		t.Fatal("expected error on cancelled context")
	}

	_, err = c.ExecuteSwap(ctx, SwapParams{})
	if err == nil {
		t.Fatal("expected error on cancelled context")
	}
}
