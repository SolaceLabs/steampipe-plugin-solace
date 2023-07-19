package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEnumVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_enum_version",
		Description: "Get a list of enumVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEnumVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEnumVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "enum_id", Type: proto.ColumnType_STRING, Description: "Enum Id."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display Name."},
			{Name: "values", Type: proto.ColumnType_JSON, Description: "Values."},
			{Name: "referenced_by_event_version_ids", Type: proto.ColumnType_STRING, Description: "Referenced by Event Version Ids."},
			{Name: "referenced_by_topic_domain_ids", Type: proto.ColumnType_STRING, Description: "Referenced by Topic Domain Ids."},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "State Id."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listEnumVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEnumVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_enumVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEnumVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		enumVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_enumVersion.listEnumVersions", "enumVersions", enumVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_enumVersion.listEnumVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range enumVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEnumVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_enumVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_enumVersion.getEnumVersion - ID", id)

	enumVersion, err := c.GetEnumVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_enumVersion.getEnumVersion", "enumVersion", fmt.Sprintf("%+v", enumVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_enumVersion.getEnumVersion", "request_error", err)
		return nil, err
	}

	return enumVersion, nil
}
