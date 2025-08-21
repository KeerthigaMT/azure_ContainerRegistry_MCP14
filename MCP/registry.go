package main

import (
	"github.com/containerregistrymanagementclient/mcp-server/config"
	"github.com/containerregistrymanagementclient/mcp-server/models"
	tools_registries "github.com/containerregistrymanagementclient/mcp-server/tools/registries"
	tools_webhooks "github.com/containerregistrymanagementclient/mcp-server/tools/webhooks"
	tools_operation "github.com/containerregistrymanagementclient/mcp-server/tools/operation"
	tools_replications "github.com/containerregistrymanagementclient/mcp-server/tools/replications"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_registries.CreateRegistries_importimageTool(cfg),
		tools_webhooks.CreateWebhooks_listTool(cfg),
		tools_webhooks.CreateWebhooks_pingTool(cfg),
		tools_registries.CreateRegistries_listusagesTool(cfg),
		tools_registries.CreateRegistries_listbyresourcegroupTool(cfg),
		tools_registries.CreateRegistries_getTool(cfg),
		tools_registries.CreateRegistries_updateTool(cfg),
		tools_registries.CreateRegistries_createTool(cfg),
		tools_registries.CreateRegistries_deleteTool(cfg),
		tools_operation.CreateOperations_listTool(cfg),
		tools_registries.CreateRegistries_listTool(cfg),
		tools_registries.CreateRegistries_listcredentialsTool(cfg),
		tools_operation.CreateRegistries_checknameavailabilityTool(cfg),
		tools_replications.CreateReplications_deleteTool(cfg),
		tools_replications.CreateReplications_getTool(cfg),
		tools_replications.CreateReplications_updateTool(cfg),
		tools_replications.CreateReplications_createTool(cfg),
		tools_registries.CreateRegistries_regeneratecredentialTool(cfg),
		tools_webhooks.CreateWebhooks_createTool(cfg),
		tools_webhooks.CreateWebhooks_deleteTool(cfg),
		tools_webhooks.CreateWebhooks_getTool(cfg),
		tools_webhooks.CreateWebhooks_updateTool(cfg),
		tools_replications.CreateReplications_listTool(cfg),
		tools_webhooks.CreateWebhooks_getcallbackconfigTool(cfg),
		tools_webhooks.CreateWebhooks_listeventsTool(cfg),
	}
}
