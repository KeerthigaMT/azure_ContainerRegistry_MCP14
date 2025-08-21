package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/containerregistrymanagementclient/mcp-server/config"
	"github.com/containerregistrymanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Replications_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		registryNameVal, ok := args["registryName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: registryName"), nil
		}
		registryName, ok := registryNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: registryName"), nil
		}
		replicationNameVal, ok := args["replicationName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: replicationName"), nil
		}
		replicationName, ok := replicationNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: replicationName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerRegistry/registries/%s/replications/%s%s", cfg.BaseURL, subscriptionId, resourceGroupName, registryName, replicationName, queryString)
		req, err := http.NewRequest("PATCH", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateReplications_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft.ContainerRegistry_registries_registryName_replications_replicationName",
		mcp.WithDescription("Updates a replication for a container registry with the specified parameters."),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("The client API version.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("The Microsoft Azure subscription ID.")),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group to which the container registry belongs.")),
		mcp.WithString("registryName", mcp.Required(), mcp.Description("The name of the container registry.")),
		mcp.WithString("replicationName", mcp.Required(), mcp.Description("The name of the replication.")),
		mcp.WithObject("replicationUpdateParameters", mcp.Required(), mcp.Description("The parameters for updating a replication.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Replications_updateHandler(cfg),
	}
}
