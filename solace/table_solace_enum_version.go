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
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "enumId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "values", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "referencedByEventVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "referencedByTopicDomainIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEnumVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEnumVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_enumVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEnumVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		enumVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_enumVersion.listEnumVersions", "enumVersions", enumVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_enumVersion.listEnumVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Trace("DEBUGGING solace_enumVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_enumVersion.getEnumVersion - ID", id)

	enumVersion, err := c.GetEnumVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_enumVersion.getEnumVersion", "enumVersion", fmt.Sprintf("%+v", enumVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_enumVersion.getEnumVersion", "request_error", err)
		return nil, err
	}

	return enumVersion, nil
}
