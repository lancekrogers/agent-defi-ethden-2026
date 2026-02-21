package payment

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// testProtocol creates a PaymentProtocol pointing to a mock HTTP server.
func testProtocol(t *testing.T, rpcHandler http.HandlerFunc) (PaymentProtocol, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(rpcHandler)
	t.Cleanup(srv.Close)

	p := NewProtocol(ProtocolConfig{
		RPCURL:        srv.URL,
		ChainID:       84532,
		WalletAddress: "0xabcdef1234567890",
		HTTPTimeout:   5 * time.Second,
	})
	return p, srv
}

// balanceHandler returns an eth_getBalance RPC response.
func balanceHandler(balanceHex string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  balanceHex,
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func TestPay_Success(t *testing.T) {
	// Handler that returns a high balance for eth_getBalance.
	callCount := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		callCount++
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  "0xde0b6b3a7640000", // 1 ETH in wei (hex)
		}
		json.NewEncoder(w).Encode(resp)
	}

	p, _ := testProtocol(t, handler)

	req := PaymentRequest{
		InvoiceID: "inv-001",
		Recipient: "0xrecipient",
		AmountWei: "0x38d7ea4c68000", // 0.001 ETH
		Network:   84532,
	}

	receipt, err := p.Pay(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if receipt == nil {
		t.Fatal("expected receipt, got nil")
	}
	if receipt.InvoiceID != "inv-001" {
		t.Errorf("expected inv-001, got %s", receipt.InvoiceID)
	}
	if receipt.ProofHeader == "" {
		t.Error("expected non-empty proof header")
	}
}

func TestPay_InsufficientFunds(t *testing.T) {
	// Return a very low balance.
	p, _ := testProtocol(t, balanceHandler("0x1")) // 1 wei

	req := PaymentRequest{
		InvoiceID: "inv-002",
		Recipient: "0xrecipient",
		AmountWei: "0xde0b6b3a7640000", // 1 ETH
		Network:   84532,
	}

	_, err := p.Pay(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for insufficient funds")
	}
	if !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("expected ErrInsufficientFunds, got %v", err)
	}
}

func TestPay_InvalidInvoice(t *testing.T) {
	tests := []struct {
		name string
		req  PaymentRequest
	}{
		{
			name: "missing InvoiceID",
			req:  PaymentRequest{Recipient: "0xr", AmountWei: "0x1"},
		},
		{
			name: "missing Recipient",
			req:  PaymentRequest{InvoiceID: "inv-1", AmountWei: "0x1"},
		},
		{
			name: "missing Amount",
			req:  PaymentRequest{InvoiceID: "inv-1", Recipient: "0xr"},
		},
		{
			name: "network mismatch",
			req:  PaymentRequest{InvoiceID: "inv-1", Recipient: "0xr", AmountWei: "0x1", Network: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := testProtocol(t, balanceHandler("0xde0b6b3a7640000"))

			_, err := p.Pay(context.Background(), tt.req)
			if err == nil {
				t.Fatal("expected error for invalid invoice")
			}
			if !errors.Is(err, ErrInvalidInvoice) {
				t.Errorf("expected ErrInvalidInvoice, got %v", err)
			}
		})
	}
}

func TestVerify_Success(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"status":      "0x1",
				"blockNumber": "0x12345",
				"from":        "0xsender",
				"to":          "0xrecipient",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}

	p, _ := testProtocol(t, handler)

	receipt, err := p.VerifyPayment(context.Background(), "inv-001", "0xtxhash")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if receipt == nil {
		t.Fatal("expected receipt, got nil")
	}
	if receipt.TxHash != "0xtxhash" {
		t.Errorf("expected 0xtxhash, got %s", receipt.TxHash)
	}
}

func TestVerify_InvalidInvoice(t *testing.T) {
	tests := []struct {
		name      string
		invoiceID string
		txHash    string
	}{
		{name: "empty invoiceID", invoiceID: "", txHash: "0x1"},
		{name: "empty txHash", invoiceID: "inv-1", txHash: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := testProtocol(t, balanceHandler("0x1"))

			_, err := p.VerifyPayment(context.Background(), tt.invoiceID, tt.txHash)
			if err == nil {
				t.Fatal("expected error")
			}
			if !errors.Is(err, ErrInvalidInvoice) {
				t.Errorf("expected ErrInvalidInvoice, got %v", err)
			}
		})
	}
}

func TestContextCancelled(t *testing.T) {
	tests := []struct {
		name string
		fn   func(ctx context.Context, p PaymentProtocol) error
	}{
		{
			name: "Pay",
			fn: func(ctx context.Context, p PaymentProtocol) error {
				_, err := p.Pay(ctx, PaymentRequest{InvoiceID: "i", Recipient: "r", AmountWei: "0x1"})
				return err
			},
		},
		{
			name: "RequestPayment",
			fn: func(ctx context.Context, p PaymentProtocol) error {
				_, err := p.RequestPayment(ctx, "0x1", "test")
				return err
			},
		},
		{
			name: "VerifyPayment",
			fn: func(ctx context.Context, p PaymentProtocol) error {
				_, err := p.VerifyPayment(ctx, "inv-1", "0xtx")
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProtocol(ProtocolConfig{
				RPCURL:        "http://203.0.113.0:9999",
				ChainID:       84532,
				WalletAddress: "0xaddr",
				HTTPTimeout:   5 * time.Second,
			})

			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			err := tt.fn(ctx, p)
			if err == nil {
				t.Fatal("expected error for cancelled context")
			}
		})
	}
}

func TestRequestPayment_Success(t *testing.T) {
	p := NewProtocol(ProtocolConfig{
		RPCURL:        "http://localhost:9999",
		ChainID:       84532,
		WalletAddress: "0xmyaddress",
	})

	invoice, err := p.RequestPayment(context.Background(), "0x38d7ea4c68000", "test compute")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if invoice == nil {
		t.Fatal("expected invoice, got nil")
	}
	if invoice.PayTo != "0xmyaddress" {
		t.Errorf("expected 0xmyaddress, got %s", invoice.PayTo)
	}
	if invoice.AmountWei != "0x38d7ea4c68000" {
		t.Errorf("expected 0x38d7ea4c68000, got %s", invoice.AmountWei)
	}
	if invoice.Network != 84532 {
		t.Errorf("expected 84532, got %d", invoice.Network)
	}
	if invoice.ExpiresAt.Before(time.Now()) {
		t.Error("invoice should not be expired immediately")
	}
}
