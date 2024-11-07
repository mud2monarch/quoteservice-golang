# quoteservice-golang

# Uniswap Telegram Quote Bot

A Telegram bot that provides Uniswap quotes and trade links. Users can request quotes with natural language (e.g. "Buy 1 ETH with USDC").

## Features
- Parse swap requests
- Fetch Uniswap quotes
- Generate trade deeplinks
- Return formatted responses

## Technical Stack
- Go 1.16+
- Telegram Bot API 
- Uniswap API

## Implementation Steps

1. **Bot Setup**
- Create Telegram bot token
- Set up Go project
- Implement basic message handling

2. **Core Functions** 
- Parse user messages
- Call Uniswap API
- Generate deeplinks
- Format responses

3. **Project Structure**
```
├── main.go
├── bot/
│   └── handler.go
├── parser/
│   └── message.go
└── uniswap/
└── client.go
```

## Example Usage
User: "Buy 1 ETH with USDC"
Bot: "2680 USDC https://app.uniswap.org/#/swap?..."

Uniswap API: https://uniswap-docs.readme.io/reference/trading-flow