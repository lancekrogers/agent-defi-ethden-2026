// Package payment implements the x402 protocol for machine-to-machine payments.
//
// x402 is a payment protocol using HTTP 402 Payment Required responses for
// micropayment negotiation between autonomous agents. An agent requests a
// resource, receives a 402 with payment details, pays via Base Sepolia, then
// retries with a payment proof header.
package payment

import (
	"errors"
	"time"
)

// Sentinel errors for payment operations.
var (
	// ErrPaymentFailed is returned when a payment transaction fails on-chain.
	ErrPaymentFailed = errors.New("payment: transaction failed on Base chain")

	// ErrInsufficientFunds is returned when the agent's wallet has insufficient
	// balance to complete the payment.
	ErrInsufficientFunds = errors.New("payment: insufficient funds in agent wallet")

	// ErrInvalidInvoice is returned when the invoice format is malformed or
	// contains invalid payment details.
	ErrInvalidInvoice = errors.New("payment: invalid invoice format or details")
)

// PaymentRequest holds the details for initiating an x402 payment.
type PaymentRequest struct {
	// InvoiceID is the unique identifier for this payment request.
	InvoiceID string

	// Recipient is the Ethereum address to receive the payment.
	Recipient string

	// AmountWei is the payment amount in wei (smallest ETH unit).
	AmountWei string

	// Token is the ERC-20 token address. Empty string means native ETH.
	Token string

	// Network is the target chain ID (e.g., 84532 for Base Sepolia).
	Network int64

	// PaymentURL is the endpoint to confirm payment after sending.
	PaymentURL string

	// Memo is an optional note attached to the payment.
	Memo string
}

// Receipt is returned after a successful x402 payment.
type Receipt struct {
	// InvoiceID is the invoice this receipt corresponds to.
	InvoiceID string

	// TxHash is the on-chain transaction hash of the payment.
	TxHash string

	// AmountWei is the amount paid in wei.
	AmountWei string

	// Sender is the Ethereum address that sent the payment.
	Sender string

	// Recipient is the Ethereum address that received the payment.
	Recipient string

	// Network is the chain ID where the payment was made.
	Network int64

	// PaidAt is when the payment transaction was confirmed.
	PaidAt time.Time

	// BlockNumber is the block at which the payment was confirmed.
	BlockNumber uint64

	// ProofHeader is the x402 payment proof header value to include in
	// subsequent requests to the resource.
	ProofHeader string
}

// Invoice represents a payment request received from a resource server
// via an HTTP 402 response.
type Invoice struct {
	// InvoiceID is the unique identifier for this payment request.
	InvoiceID string `json:"invoice_id"`

	// PayTo is the Ethereum address to pay.
	PayTo string `json:"pay_to"`

	// AmountWei is the required payment amount in wei.
	AmountWei string `json:"amount_wei"`

	// Token is the ERC-20 token address. Empty means native ETH.
	Token string `json:"token,omitempty"`

	// Network is the required chain ID.
	Network int64 `json:"network"`

	// Description describes what the payment is for.
	Description string `json:"description,omitempty"`

	// ExpiresAt is when this invoice expires.
	ExpiresAt time.Time `json:"expires_at"`

	// CallbackURL is where to POST the payment proof.
	CallbackURL string `json:"callback_url"`
}
