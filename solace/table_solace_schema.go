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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "application_domain_id", Type: proto.ColumnType_STRING, Description: "Application Domain Id."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name."},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: "Shared?."},
			{Name: "schema_type", Type: proto.ColumnType_STRING, Description: "Schema Type."},
			{Name: "number_of_versions", Type: proto.ColumnType_STRING, Description: "Number of Versions."},
			{Name: "event_version_ref_count", Type: proto.ColumnType_STRING, Description: "Event Version Ref Count."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listSchemas", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_schema.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewSchemaListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		schemas, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_schema.listSchemas", "schemas", schemas)
		if err != nil {
			plugin.Logger(ctx).Error("solace_schema.listSchemas", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_schema.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_schema.getSchema - ID", id)

	schema, err := c.GetSchema(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_schema.getSchema", "schema", fmt.Sprintf("%+v", schema))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_schema.getSchema", "request_error", err)
		return nil, err
	}

	return schema, nil
}
