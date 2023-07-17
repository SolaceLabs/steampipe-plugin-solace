package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEnum(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_enum",
		Description: "Get a list of enums that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEnums,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEnum,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "application_domain_id", Type: proto.ColumnType_STRING, Description: "Application Domain Id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: "Shared?"},
			{Name: "number_of_versions", Type: proto.ColumnType_STRING, Description: "Number of Versions"},
			{Name: "event_version_ref_count", Type: proto.ColumnType_STRING, Description: "Event Version Ref Count"},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEnums(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEnums", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_enum.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEnumListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		enums, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_enum.listEnums", "enums", enums)
		if err != nil {
			plugin.Logger(ctx).Error("solace_enum.listEnums", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range enums {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEnum(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_enum.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")

	enum, err := c.GetEnum(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_enum.getEnum", "enum", fmt.Sprintf("%+v", enum))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_enum.getEnum", "request_error", err)
		return nil, err
	}

	return enum, nil
}
