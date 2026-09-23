# Arbitrage Bot

A Go-based bot for monitoring marketplaces and finding arbitrage opportunities between equivalent offers.

## What it does

The bot collects listings from different marketplaces and extracts their key data:

**Marketplace → Parser → Offers → Matching → Arbitrage**

The long-term goal is to automatically detect situations where the same product or service can be bought on one marketplace and sold on another at a profit after marketplace fees.

## Marketplaces

| Marketplace | Status            | Method             |
| ----------- | ----------------- | ------------------ |
| FunPay      | 🟡 In development | Browser automation |
| Playerok    | 🟡 In development | Browser automation |
| GGsel       | ⚪ Planned         | —                  |

## Current Data

### FunPay

Currently extracted:

* Service
* Seller
* Amount
* Type
* Price
* Marketplace

### Playerok

Currently extracted:

* Product name
* Price
* Marketplace

The parsers run concurrently using Go goroutines and communicate with the main program through channels.

## Tech Stack

**Core**

* Go
* chromedp

**Planned**

* PostgreSQL
* Telegram Bot API
* Docker
* Additional HTTP-based integrations where possible

## Project Flow

```text
                    ┌──────────┐
                    │ FunPay   │
                    └────┬─────┘
                         │
                         ▼
                    ┌──────────┐
                    │  Parser  │
                    └────┬─────┘
                         │
                         │
                         ▼
┌──────────┐        ┌──────────┐
│ Playerok │───────▶│  Offers  │
└────┬─────┘        └────┬─────┘
     │                   │
     ▼                   ▼
┌──────────┐       ┌────────────┐
│  Parser  │       │   Matcher  │
└──────────┘       └──────┬─────┘
                          │
                          ▼
                   ┌─────────────┐
                   │  Arbitrage  │
                   │  Calculator │
                   └──────┬──────┘
                          │
                          ▼
                   ┌─────────────┐
                   │ PostgreSQL  │
                   └──────┬──────┘
                          │
                          ▼
                   ┌─────────────┐
                   │  Telegram   │
                   └─────────────┘
```

## Roadmap

### 1. Data collection

* [x] FunPay parser
* [x] Playerok parser
* [ ] Improve dynamic page handling
* [ ] Replace fixed delays with proper DOM conditions
* [ ] Investigate direct HTTP requests
* [ ] Add more marketplaces

### 2. Data processing

* [ ] Normalize product names
* [ ] Normalize prices and currencies
* [ ] Match equivalent offers
* [ ] Handle different quantities
* [ ] Account for marketplace fees

### 3. Arbitrage engine

* [ ] Calculate potential profit
* [ ] Calculate profit percentage
* [ ] Filter unprofitable opportunities
* [ ] Track opportunities over time

### 4. Storage

* [ ] PostgreSQL integration
* [ ] Store offers
* [ ] Store price history
* [ ] Store arbitrage opportunities

### 5. Notifications

* [ ] Telegram bot
* [ ] Arbitrage alerts
* [ ] Configurable minimum profit

### 6. Infrastructure

* [ ] Environment-based configuration
* [ ] Logging
* [ ] Tests
* [ ] Docker

## Project Structure

```text
arbitrage-bot/
├── main.go
├── funpay.go
├── playerok.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

## Status

🚧 **Early development**

The project is currently focused on building reliable marketplace parsers and understanding the data provided by each marketplace before implementing the arbitrage engine.
