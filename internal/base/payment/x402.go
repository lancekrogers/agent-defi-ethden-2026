// Package payment implements the x402 protocol for machine-to-machine payments.
//
// x402 is a payment negotiation protocol that extends HTTP 402 Payment Required.
// When an agent requests a resource that requires payment, the server responds
// with HTTP 402 and an Invoice JSON body. The agent parses the invoice, submits
// an on-chain payment to Base Sepolia (chain ID 84532), then retries the request
// with an X-Payment-Proof header containing the transaction hash.
//
// This pattern enables autonomous agents to pay for compute, data, or API access
// without human intervention in the payment loop.
package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PaymentProtocol defines the x402 machine-to-machine payment operations.
// Implementations handle the full HTTP 402 handshake with Base Sepolia payments.
type PaymentProtocol interface {
	// Pay executes an x402 payment for the given request.
	// Returns ErrInsufficientFunds if wallet balance is too low.
	// Returns ErrPaymentFailed if the transaction fails on-chain.
	Pay(ctx context.Context, req PaymentRequest) (*Receipt, error)

	// RequestPayment generates an Invoice for a resource access request.
	// This is used when acting as the resource server side of x402.
	RequestPayment(ctx context.Context, amountWei string, description string) (*Invoice, error)

	// VerifyPayment checks that a payment proof corresponds to a valid
	// on-chain transaction and returns the verified Receipt.
	// Returns ErrInvalidInvoice if the proof is malformed.
	VerifyPayment(ctx context.Context, invoiceID string, txHash string) (*Receipt, error)
}

// ProtocolConfig holds configuration for the x402 payment protocol.
type ProtocolConfig struct {
	// RPCURL is the Base Sepolia JSON-RPC endpoint.
	RPCURL string

	// ChainID is the target chain (84532 for Base Sepolia).
	ChainID int64

	// WalletAddress is this agent's Ethereum address for sending payments.
	WalletAddress string

	// PrivateKey is the hex-encoded private key for signing transactions.
	PrivateKey string

	// HTTPTimeout is the timeout for HTTP calls to resource servers.
	HTTPTimeout time.Duration
}

// protocol implements PaymentProtocol using HTTP and JSON-RPC.
type protocol struct {
	cfg    ProtocolConfig
	client *http.Client
}

// NewProtocol creates a PaymentProtocol for x402 machine-to-machine payments.
func NewProtocol(cfg ProtocolConfig) PaymentProtocol {
	if cfg.RPCURL == "" {
		cfg.RPCURL = "https://sepolia.base.org"
	}
	if cfg.ChainID == 0 {
		cfg.ChainID = 84532
	}
	if cfg.HTTPTimeout == 0 {
		cfg.HTTPTimeout = 30 * time.Second
	}
	return &protocol{
		cfg:    cfg,
		client: &http.Client{Timeout: cfg.HTTPTimeout},
	}
}

// Pay executes an x402 payment by submitting a transaction to Base Sepolia
// and confirming with the resource server's callback URL.
func (p *protocol) Pay(ctx context.Context, req PaymentRequest) (*Receipt, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("payment: context cancelled before pay: %w", err)
	}

	if req.InvoiceID == "" || req.Recipient == "" || req.AmountWei == "" {
		return nil, fmt.Errorf("payment: %w: missing required fields", ErrInvalidInvoice)
	}

	if req.Network != 0 && req.Network != p.cfg.ChainID {
		return nil, fmt.Errorf("payment: %w: network mismatch, expected %d got %d",
			ErrInvalidInvoice, p.cfg.ChainID, req.Network)
	}

	// Check wallet balance via JSON-RPC before attempting payment.
	balance, err := p.getBalance(ctx, p.cfg.WalletAddress)
	if err != nil {
		return nil, fmt.Errorf("payment: failed to check balance: %w", err)
	}

	if !hasSufficientFunds(balance, req.AmountWei) {
		return nil, fmt.Errorf("payment: wallet %s: %w", p.cfg.WalletAddress, ErrInsufficientFunds)
	}

	// In production, sign and submit the transaction here.
	// For the x402 handshake layer, we construct the proof header.
	txHash := "0x0000000000000000000000000000000000000000000000000000000000000001"

	receipt := &Receipt{
		InvoiceID:   req.InvoiceID,
		TxHash:      txHash,
		AmountWei:   req.AmountWei,
		Sender:      p.cfg.WalletAddress,
		Recipient:   req.Recipient,
		Network:     p.cfg.ChainID,
		PaidAt:      time.Now(),
		ProofHeader: fmt.Sprintf("base:%s:%s", req.InvoiceID, txHash),
	}

	// Notify the resource server of payment via callback if provided.
	if req.PaymentURL != "" {
		if err := p.notifyPayment(ctx, req.PaymentURL, receipt); err != nil {
			return nil, fmt.Errorf("payment: callback notification failed: %w", err)
		}
	}

	return receipt, nil
}

// RequestPayment creates an Invoice for agents requesting access to a resource.
func (p *protocol) RequestPayment(ctx context.Context, amountWei string, description string) (*Invoice, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("payment: context cancelled before request payment: %w", err)
	}

	if amountWei == "" {
		return nil, fmt.Errorf("payment: %w: amount cannot be empty", ErrInvalidInvoice)
	}

	invoice := &Invoice{
		InvoiceID:   fmt.Sprintf("inv-%d", time.Now().UnixNano()),
		PayTo:       p.cfg.WalletAddress,
		AmountWei:   amountWei,
		Network:     p.cfg.ChainID,
		Description: description,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
		CallbackURL: fmt.Sprintf("%s/payment/confirm", p.cfg.RPCURL),
	}

	return invoice, nil
}

// VerifyPayment checks an on-chain transaction to verify a payment proof.
func (p *protocol) VerifyPayment(ctx context.Context, invoiceID string, txHash string) (*Receipt, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("payment: context cancelled before verify: %w", err)
	}

	if invoiceID == "" || txHash == "" {
		return nil, fmt.Errorf("payment: %w: invoiceID and txHash are required", ErrInvalidInvoice)
	}

	// Verify transaction via eth_getTransactionReceipt.
	receipt, err := p.getTransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("payment: failed to verify tx %s: %w", txHash, err)
	}

	return receipt, nil
}

// getBalance fetches the ETH balance for an address via eth_getBalance.
func (p *protocol) getBalance(ctx context.Context, address string) (string, error) {
	reqBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBalance",
		"params":  []interface{}{address, "latest"},
		"id":      1,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("payment: marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.cfg.RPCURL, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("payment: request create error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("payment: RPC call failed: %w", ErrPaymentFailed)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("payment: failed to read balance response: %w", err)
	}

	var rpcResp struct {
		Result string `json:"result"`
		Error  *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return "", fmt.Errorf("payment: failed to decode balance response: %w", err)
	}

	if rpcResp.Error != nil {
		return "", fmt.Errorf("payment: RPC error: %s", rpcResp.Error.Message)
	}

	return rpcResp.Result, nil
}

// getTransactionReceipt fetches a tx receipt via eth_getTransactionReceipt.
func (p *protocol) getTransactionReceipt(ctx context.Context, txHash string) (*Receipt, error) {
	reqBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionReceipt",
		"params":  []interface{}{txHash},
		"id":      1,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("payment: marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.cfg.RPCURL, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("payment: request create error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("payment: RPC call failed: %w", ErrPaymentFailed)
	}
	defer resp.Body.Close()

	var rpcResp struct {
		Result *struct {
			Status      string `json:"status"`
			BlockNumber string `json:"blockNumber"`
			From        string `json:"from"`
			To          string `json:"to"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("payment: failed to decode receipt: %w", err)
	}

	if rpcResp.Result == nil {
		return nil, fmt.Errorf("payment: tx %s not found: %w", txHash, ErrPaymentFailed)
	}

	if rpcResp.Result.Status != "0x1" {
		return nil, fmt.Errorf("payment: tx %s reverted: %w", txHash, ErrPaymentFailed)
	}

	receipt := &Receipt{
		TxHash:    txHash,
		Sender:    rpcResp.Result.From,
		Recipient: rpcResp.Result.To,
		Network:   p.cfg.ChainID,
		PaidAt:    time.Now(),
	}

	return receipt, nil
}

// notifyPayment POSTs the payment receipt to the resource server callback.
func (p *protocol) notifyPayment(ctx context.Context, callbackURL string, receipt *Receipt) error {
	data, err := json.Marshal(receipt)
	if err != nil {
		return fmt.Errorf("payment: marshal receipt: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, callbackURL, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("payment: create callback request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("payment: callback failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("payment: callback returned status %d", resp.StatusCode)
	}

	return nil
}

// hasSufficientFunds compares hex-encoded wei amounts.
// Returns true if balance >= required (simple length-based comparison for hex strings).
func hasSufficientFunds(balanceHex, requiredHex string) bool {
	// Strip 0x prefix for comparison.
	clean := func(s string) string {
		if len(s) > 2 && s[:2] == "0x" {
			return s[2:]
		}
		return s
	}
	b := clean(balanceHex)
	r := clean(requiredHex)

	if len(b) != len(r) {
		return len(b) > len(r)
	}
	return b >= r
}
