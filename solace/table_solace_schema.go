package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableSchema(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_schema",
		Description: "Get a list of schemas that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listSchemas,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSchema,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationDomainId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "schemaType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numberOfVersions", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventVersionRefCount", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listSchemas", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_schema.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewSchemaListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		schemas, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_schema.listSchemas", "schemas", schemas)
		if err != nil {
			plugin.Logger(ctx).Error("solace_schema.listSchemas", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range schemas {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getSchema(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_schema.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_schema.getSchema - ID", id)

	schema, err := c.GetSchema(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_schema.getSchema", "schema", fmt.Sprintf("%+v", schema))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_schema.getSchema", "request_error", err)
		return nil, err
	}

	return schema, nil
}
