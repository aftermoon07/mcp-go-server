package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"HelloMCPServer",
		"1.0.0",
	)

	// Define a simple "echo" tool
	echoTool := mcp.NewTool("echo",
		mcp.WithDescription("Echoes the input text back to the user"),
		mcp.WithString("text", mcp.Required(), mcp.Description("The text to echo")),
	)

	// Add the tool and its handler to the server
	s.AddTool(echoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("Arguments must be a map"), nil
		}

		text, ok := args["text"].(string)
		if !ok {
			return mcp.NewToolResultError("The 'text' argument must be a string"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Echo from Go MCP: %s", text)), nil
	})

	// Start the server using stdio
	// This is the standard way MCP servers communicate with hosts
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
