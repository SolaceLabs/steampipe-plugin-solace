package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_configuration",
		Description: "Get a list of configurations that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listConfigurations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getConfiguration,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contextType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "contextId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "configurationTypeId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "entityType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "entityId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listConfigurations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listConfigurations", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_configuration.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewConfigurationListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		configurations, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_configuration.listConfigurations", "configurations", configurations)
		if err != nil {
			plugin.Logger(ctx).Error("solace_configuration.listConfigurations", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range configurations {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getConfiguration(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_configuration.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_configuration.getConfiguration - ID", id)

	configuration, err := c.GetConfiguration(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_configuration.getConfiguration", "configuration", fmt.Sprintf("%+v", configuration))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_configuration.getConfiguration", "request_error", err)
		return nil, err
	}

	return configuration, nil
}
