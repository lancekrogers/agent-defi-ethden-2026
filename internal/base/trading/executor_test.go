package trading

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// testExecutor creates an executor pointing to a mock HTTP server.
func testExecutor(t *testing.T, handler http.HandlerFunc) (TradeExecutor, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	exec := NewExecutor(ExecutorConfig{
		RPCURL:           srv.URL,
		ChainID:          84532,
		WalletAddress:    "0xagentaddress",
		DEXRouterAddress: "0xrouteraddress",
		HTTPTimeout:      5 * time.Second,
	})
	return exec, srv
}

// rpcResultHandler returns an HTTP handler serving a fixed JSON-RPC result.
func rpcResultHandler(result interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resultData, _ := json.Marshal(result)
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  json.RawMessage(resultData),
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func TestExecute_Success(t *testing.T) {
	exec, _ := testExecutor(t, rpcResultHandler("0x1234567"))

	trade := Trade{
		TokenIn:      "0xusdc",
		TokenOut:     "0xweth",
		AmountIn:     "0x1000",
		MinAmountOut: "0x100",
		Deadline:     time.Now().Add(5 * time.Minute),
		Signal: Signal{
			Type:       SignalBuy,
			Confidence: 0.8,
		},
	}

	result, err := exec.Execute(context.Background(), trade)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
	if result.TxHash == "" {
		t.Error("expected non-empty tx hash")
	}
	if result.ExecutedAt.IsZero() {
		t.Error("expected non-zero executed at time")
	}
}

func TestExecute_MissingTokens(t *testing.T) {
	exec, _ := testExecutor(t, rpcResultHandler("0x1"))

	trade := Trade{
		AmountIn: "0x1000",
	}

	_, err := exec.Execute(context.Background(), trade)
	if err == nil {
		t.Fatal("expected error for missing token addresses")
	}
}

func TestExecute_ContextCancelled(t *testing.T) {
	exec := NewExecutor(ExecutorConfig{
		RPCURL:        "http://203.0.113.0:9999",
		ChainID:       84532,
		WalletAddress: "0xagent",
		HTTPTimeout:   5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	trade := Trade{
		TokenIn:  "0xusdc",
		TokenOut: "0xweth",
	}

	_, err := exec.Execute(ctx, trade)
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}

func TestGetBalance_ETH(t *testing.T) {
	// Return a hex ETH balance.
	exec, _ := testExecutor(t, rpcResultHandler("0xde0b6b3a7640000"))

	balance, err := exec.GetBalance(context.Background(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if balance == nil {
		t.Fatal("expected balance, got nil")
	}
	if balance.AmountWei != "0xde0b6b3a7640000" {
		t.Errorf("expected 0xde0b6b3a7640000, got %s", balance.AmountWei)
	}
	if balance.TokenAddress != "" {
		t.Errorf("expected empty token address for ETH, got %s", balance.TokenAddress)
	}
	if balance.UpdatedAt.IsZero() {
		t.Error("expected non-zero updated at")
	}
}

func TestGetBalance_ERC20(t *testing.T) {
	exec, _ := testExecutor(t, rpcResultHandler("0x1000"))

	balance, err := exec.GetBalance(context.Background(), "0xtokenaddress")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if balance.TokenAddress != "0xtokenaddress" {
		t.Errorf("expected 0xtokenaddress, got %s", balance.TokenAddress)
	}
}

func TestGetBalance_ContextCancelled(t *testing.T) {
	exec := NewExecutor(ExecutorConfig{
		RPCURL:        "http://203.0.113.0:9999",
		WalletAddress: "0xagent",
		HTTPTimeout:   5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := exec.GetBalance(ctx, "")
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}

func TestGetMarketState_Success(t *testing.T) {
	exec, _ := testExecutor(t, rpcResultHandler("0x1234567"))

	state, err := exec.GetMarketState(context.Background(), "0xusdc", "0xweth")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state == nil {
		t.Fatal("expected market state, got nil")
	}
	if state.Price <= 0 {
		t.Error("expected positive price")
	}
	if state.MovingAverage <= 0 {
		t.Error("expected positive moving average")
	}
	if state.FetchedAt.IsZero() {
		t.Error("expected non-zero fetched at")
	}
	if state.TokenIn != "0xusdc" {
		t.Errorf("expected 0xusdc, got %s", state.TokenIn)
	}
	if state.TokenOut != "0xweth" {
		t.Errorf("expected 0xweth, got %s", state.TokenOut)
	}
}

func TestGetMarketState_ContextCancelled(t *testing.T) {
	exec := NewExecutor(ExecutorConfig{
		RPCURL:      "http://203.0.113.0:9999",
		HTTPTimeout: 5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := exec.GetMarketState(ctx, "0xusdc", "0xweth")
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}
