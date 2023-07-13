package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCustomAttributeDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_custom_attribute_definition",
		Description: "Get a list of custom attribute definitions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listCustomAttributeDefinitions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCustomAttributeDefinition,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "valueType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "scope", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "associatedEntityTypes", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "associatedEntities", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listCustomAttributeDefinitions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listCustomAttributeDefinitions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_customAttributeDefinition.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewCustomAttributeDefinitionListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		customAttributeDefinitions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_customAttributeDefinition.listCustomAttributeDefinitions", "customAttributeDefinitions", customAttributeDefinitions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_customAttributeDefinition.listCustomAttributeDefinitions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range customAttributeDefinitions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getCustomAttributeDefinition(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_customAttributeDefinition.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_customAttributeDefinition.getCustomAttributeDefinition - ID", id)

	customAttributeDefinition, err := c.GetCustomAttributeDefinition(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_customAttributeDefinition.getCustomAttributeDefinition", "customAttributeDefinition", fmt.Sprintf("%+v", customAttributeDefinition))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_customAttributeDefinition.getCustomAttributeDefinition", "request_error", err)
		return nil, err
	}

	return customAttributeDefinition, nil
}
