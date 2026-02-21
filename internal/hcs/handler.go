package hcs

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"
)

// Transport abstracts HCS topic operations for testability.
// In production this wraps the Hedera SDK; in tests it uses a mock.
type Transport interface {
	// Publish sends raw bytes to an HCS topic.
	Publish(ctx context.Context, topicID string, data []byte) error

	// Subscribe starts receiving messages from an HCS topic.
	// Messages are delivered to the returned channel until ctx is cancelled.
	Subscribe(ctx context.Context, topicID string) (<-chan []byte, <-chan error)
}

// HandlerConfig holds configuration for the DeFi agent HCS handler.
type HandlerConfig struct {
	// Transport is the HCS transport implementation.
	Transport Transport

	// TaskTopicID is the HCS topic for receiving task assignments.
	TaskTopicID string

	// ResultTopicID is the HCS topic for publishing results and reports.
	ResultTopicID string

	// AgentID is this agent's unique identifier.
	AgentID string
}

// Handler manages HCS subscriptions and publishing for the DeFi agent.
// It subscribes to task assignments and publishes P&L reports, strategy
// updates, health status, and task results.
type Handler struct {
	cfg    HandlerConfig
	seqNum atomic.Uint64
	taskCh chan TaskAssignment
}

// NewHandler creates an HCS handler for the DeFi agent.
func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{
		cfg:    cfg,
		taskCh: make(chan TaskAssignment, 16),
	}
}

// Tasks returns a read-only channel of incoming task assignments.
func (h *Handler) Tasks() <-chan TaskAssignment {
	return h.taskCh
}

// StartSubscription begins listening for task assignments on HCS.
// It runs until the context is cancelled. Malformed messages are skipped.
func (h *Handler) StartSubscription(ctx context.Context) error {
	msgCh, errCh := h.cfg.Transport.Subscribe(ctx, h.cfg.TaskTopicID)
	if msgCh == nil {
		return ErrSubscriptionFailed
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("hcs: subscription error: %w", ErrSubscriptionFailed)
			}
		case data, ok := <-msgCh:
			if !ok {
				return nil
			}
			h.processMessage(ctx, data)
		}
	}
}

func (h *Handler) processMessage(ctx context.Context, data []byte) {
	env, err := UnmarshalEnvelope(data)
	if err != nil {
		return // skip malformed messages
	}

	if env.Type != MessageTypeTaskAssignment {
		return // skip non-task messages
	}

	// Filter: only accept messages addressed to us or broadcast.
	if env.Recipient != "" && env.Recipient != h.cfg.AgentID {
		return
	}

	var task TaskAssignment
	if err := json.Unmarshal(env.Payload, &task); err != nil {
		return // skip messages with invalid payload
	}

	select {
	case h.taskCh <- task:
	case <-ctx.Done():
	}
}

// PublishPnL sends a P&L report to the result topic via HCS.
func (h *Handler) PublishPnL(ctx context.Context, report PnLReportMessage) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("hcs: context cancelled before publish pnl: %w", err)
	}

	return h.publish(ctx, MessageTypePnLReport, report)
}

// PublishStrategyUpdate broadcasts a strategy change to the result topic.
func (h *Handler) PublishStrategyUpdate(ctx context.Context, update StrategyUpdate) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("hcs: context cancelled before publish strategy update: %w", err)
	}

	return h.publish(ctx, MessageTypeStrategyUpdate, update)
}

// PublishHealth sends a health status heartbeat to the result topic.
func (h *Handler) PublishHealth(ctx context.Context, status HealthStatus) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("hcs: context cancelled before publish health: %w", err)
	}

	return h.publish(ctx, MessageTypeHeartbeat, status)
}

// PublishResult sends a task result back to the coordinator via HCS.
func (h *Handler) PublishResult(ctx context.Context, result TaskResult) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("hcs: context cancelled before publish result: %w", err)
	}

	return h.publish(ctx, MessageTypeTaskResult, result)
}

// publish is the internal helper that wraps a payload in an Envelope and publishes it.
func (h *Handler) publish(ctx context.Context, msgType MessageType, payload interface{}) error {
	payloadData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("hcs: failed to marshal payload: %w", err)
	}

	env := Envelope{
		Type:        msgType,
		Sender:      h.cfg.AgentID,
		SequenceNum: h.seqNum.Add(1),
		Timestamp:   time.Now(),
		Payload:     payloadData,
	}

	data, err := env.Marshal()
	if err != nil {
		return fmt.Errorf("hcs: failed to marshal envelope: %w", err)
	}

	if err := h.cfg.Transport.Publish(ctx, h.cfg.ResultTopicID, data); err != nil {
		return fmt.Errorf("hcs: failed to publish %s message: %w", msgType, ErrPublishFailed)
	}

	return nil
}
