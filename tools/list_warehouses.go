package tools

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/mark3labs/mcp-go/mcp"
)

// ListWarehouses retrieves all SQL warehouses from the Databricks workspace
// and returns them as a JSON string.
func ListWarehouses(ctx context.Context, _ mcp.CallToolRequest) (interface{}, error) {
	w, err := WorkspaceClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
}
