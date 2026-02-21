# agent-defi

Autonomous DeFi trading agent for the Base Sepolia chain.

Part of the [ETHDenver 2026 Agent Economy](../README.md) submission.

## Overview

Executes mean reversion trading strategies on Base Sepolia DEX via go-ethereum. Registers on-chain identity via ERC-8004, pays for compute via x402 payment protocol, and attributes all transactions via ERC-8021 builder codes. Reports P&L and health status to the coordinator agent via Hedera Consensus Service (HCS).

## Built with Obedience Corp

This project is part of an [Obedience Corp](https://obediencecorp.com) campaign — built and planned using **camp** (campaign management) and **fest** (festival methodology). This repository, its git history, and the planning artifacts in `festivals/` are a live example of these tools in action.

The agent connects to the **obey daemon** for task coordination and event routing via `OBEY_DAEMON_SOCKET`.

## System Context

```
                    ┌─────────────┐
           tasks    │ Coordinator │    tasks
          ┌────────>│  (Hedera)   │<────────┐
          │         └─────────────┘         │
          │               │                 │
          │          assignments             │
          │               │                 │
    ┌─────┴─────┐         │         ┌───────┴──────┐
    │ Inference │         │         │  DeFi Agent  │ <-- you are here
    │   (0G)    │         └────────>│   (Base)     │
    └───────────┘                   └──────────────┘
```

## Quick Start

```bash
cp .env.example .env   # fill in Hedera + daemon values
just build
just run
```

## Prerequisites

- Go 1.24+
- Hedera testnet account ([portal.hedera.com](https://portal.hedera.com))
- Base Sepolia RPC endpoint (default: https://sepolia.base.org)

## Configuration

| Variable | Description |
|----------|-------------|
| `HEDERA_ACCOUNT_ID` | Hedera testnet account (0.0.xxx) |
| `HEDERA_PRIVATE_KEY` | Hedera private key |
| `OBEY_DAEMON_SOCKET` | Path to obey daemon Unix socket |

Trading parameters (wallet address, private key, token pair, DEX router, ERC-8004 contract, builder code) are configured via `DEFI_*` env vars -- see `internal/agent/config.go` for the full list.

## Project Structure

```
cmd/agent-defi/            Entry point, dependency wiring
internal/
  agent/                   Agent lifecycle, config, goroutine orchestration
  base/
    attribution/           ERC-8021 builder code encoder/decoder
    identity/              ERC-8004 on-chain identity registration
    payment/               x402 machine-to-machine payment protocol
    trading/               Mean reversion strategy, trade executor, P&L tracker
  hcs/                     HCS publish/subscribe transport (Hiero SDK)
```

## Development

```bash
just build      # Build binary to bin/
just run        # Run the agent
just test       # Run tests
just lint       # golangci-lint
just fmt        # gofmt
just clean      # Remove build artifacts
```

## Architecture

`main.go` wires all dependencies via constructor injection: identity registry, payment protocol, attribution encoder, trading strategy + executor, P&L tracker, and HCS handler. The agent spawns concurrent goroutines for the trading loop, P&L reporting, and health heartbeats, all managed under a signal-aware context for clean shutdown on SIGINT/SIGTERM.

The trading pipeline per cycle: fetch market state from DEX, evaluate mean reversion signal, execute trade if actionable, record result in P&L tracker, and publish to coordinator via HCS.

## License

MIT
