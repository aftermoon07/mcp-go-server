# mcp-go-server

A simple implementation of a Model Context Protocol (MCP) server written in Go.

## Overview
This project serves as a demonstration of both **Go** and the **Model Context Protocol (MCP)**. It creates a local standard IO (stdio) server that can be integrated into LLM agents to provide external tool access.

It includes a simple `echo` tool that echoes text back to the caller, highlighting how to define JSON schema inputs and handle tool execution securely in Go.

## Prerequisites
- [Go](https://go.dev/) 1.20+

## Setup & Build

1. Clone the repository and initialize the module (if needed):
```bash
go mod tidy
```

2. Build the server:
```bash
go build -o mcp-server
```

## Testing Locally

You can test the stdio server directly via terminal by passing JSON-RPC commands:

### List Tools
```bash
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | ./mcp-server
```

### Call the Echo Tool
```bash
echo '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"echo","arguments":{"text":"Hello World!"}}}' | ./mcp-server
```

## Adding to an AI Agent

To use this server with an MCP-compatible agent (like Claude Desktop or an Anthropic/Gemini MCP client), configure the client to execute the binary:

```json
{
  "mcpServers": {
    "go-echo-server": {
      "command": "/absolute/path/to/mcp-server",
      "args": []
    }
  }
}
```
