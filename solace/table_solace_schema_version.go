package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableSchemaVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_schema_version",
		Description: "Get a list of schemaVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listSchemaVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSchemaVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "schemaId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "content", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "referencedByEventVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "referencedBySchemaVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "schemaVersionReferences", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listSchemaVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listSchemaVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_schemaVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewSchemaVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		schemaVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_schemaVersion.listSchemaVersions", "schemaVersions", schemaVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_schemaVersion.listSchemaVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range schemaVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getSchemaVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_schemaVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_schemaVersion.getSchemaVersion - ID", id)

	schemaVersion, err := c.GetSchemaVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_schemaVersion.getSchemaVersion", "schemaVersion", fmt.Sprintf("%+v", schemaVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_schemaVersion.getSchemaVersion", "request_error", err)
		return nil, err
	}

	return schemaVersion, nil
}
