package strategy

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// ObeyClient implements LLMClient using the obey daemon session system.
// It creates a persistent session and sends prompts through the daemon
// rather than calling AI provider APIs directly.
type ObeyClient struct {
	Socket    string // daemon gRPC socket path (default: /tmp/obey.sock)
	Campaign  string // campaign name
	Provider  string // AI provider (e.g., "claude-code")
	Model     string // model name (e.g., "claude-sonnet-4-6")
	Festival  string // festival ID (optional)
	SessionID string // reused across calls once created
}

// Complete sends a prompt to an obey session and returns the response.
func (c *ObeyClient) Complete(ctx context.Context, prompt string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", fmt.Errorf("obey: context cancelled: %w", err)
	}

	// Create session on first call
	if c.SessionID == "" {
		id, err := c.createSession(ctx)
		if err != nil {
			return "", fmt.Errorf("obey: create session: %w", err)
		}
		c.SessionID = id
	}

	return c.sendMessage(ctx, prompt)
}

func (c *ObeyClient) createSession(ctx context.Context) (string, error) {
	args := []string{
		"session", "create",
		"--socket", c.socket(),
		"--campaign", c.Campaign,
		"--provider", c.provider(),
		"--model", c.model(),
		"--agent", "vault-trader",
	}
	if c.Festival != "" {
		args = append(args, "--festival", c.Festival)
	}

	cmd := exec.CommandContext(ctx, "obey", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("obey session create failed: %w: %s", err, stderr.String())
	}

	// Parse session ID from output like "Session: <uuid>"
	for _, line := range strings.Split(stdout.String(), "\n") {
		if strings.HasPrefix(line, "Session: ") {
			return strings.TrimPrefix(line, "Session: "), nil
		}
	}

	return "", fmt.Errorf("obey: could not parse session ID from: %s", stdout.String())
}

func (c *ObeyClient) sendMessage(ctx context.Context, message string) (string, error) {
	cmd := exec.CommandContext(ctx, "obey",
		"session", "send",
		"--socket", c.socket(),
		c.SessionID,
		message,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("obey session send failed: %w: %s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

func (c *ObeyClient) socket() string {
	if c.Socket != "" {
		return c.Socket
	}
	return "/tmp/obey.sock"
}

func (c *ObeyClient) provider() string {
	if c.Provider != "" {
		return c.Provider
	}
	return "claude-code"
}

func (c *ObeyClient) model() string {
	if c.Model != "" {
		return c.Model
	}
	return "claude-sonnet-4-6"
}
