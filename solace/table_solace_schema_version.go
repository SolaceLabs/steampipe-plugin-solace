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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "schema_id", Type: proto.ColumnType_STRING, Description: "Schema Id."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display Name."},
			{Name: "content", Type: proto.ColumnType_STRING, Description: "Content."},
			{Name: "referenced_by_event_version_ids", Type: proto.ColumnType_STRING, Description: "Referenced by Event Version Ids."},
			{Name: "referenced_by_schema_version_ids", Type: proto.ColumnType_STRING, Description: "Referenced by Schema Version Ids."},
			{Name: "schema_version_references", Type: proto.ColumnType_JSON, Description: "Schema Version References."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "State Id."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listSchemaVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listSchemaVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_schemaVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewSchemaVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		schemaVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_schemaVersion.listSchemaVersions", "schemaVersions", schemaVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_schemaVersion.listSchemaVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_schemaVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_schemaVersion.getSchemaVersion - ID", id)

	schemaVersion, err := c.GetSchemaVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_schemaVersion.getSchemaVersion", "schemaVersion", fmt.Sprintf("%+v", schemaVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_schemaVersion.getSchemaVersion", "request_error", err)
		return nil, err
	}

	return schemaVersion, nil
}
