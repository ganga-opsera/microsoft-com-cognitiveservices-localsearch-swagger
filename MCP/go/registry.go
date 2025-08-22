package main

import (
	"github.com/local-search-client/mcp-server/config"
	"github.com/local-search-client/mcp-server/models"
	tools_localsearch "github.com/local-search-client/mcp-server/tools/localsearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_localsearch.CreateLocal_searchTool(cfg),
	}
}
