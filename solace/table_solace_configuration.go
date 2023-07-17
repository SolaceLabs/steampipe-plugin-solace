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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "context_type", Type: proto.ColumnType_STRING, Description: "Context Type"},
			{Name: "context_id", Type: proto.ColumnType_STRING, Description: "Context Id"},
			{Name: "configuration_type_id", Type: proto.ColumnType_STRING, Description: "Configuration Type Id"},
			{Name: "entity_type", Type: proto.ColumnType_STRING, Description: "Entity Type"},
			{Name: "entity_id", Type: proto.ColumnType_STRING, Description: "Entity Id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listConfigurations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listConfigurations", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_configuration.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewConfigurationListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		configurations, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_configuration.listConfigurations", "configurations", configurations)
		if err != nil {
			plugin.Logger(ctx).Error("solace_configuration.listConfigurations", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_configuration.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_configuration.getConfiguration - ID", id)

	configuration, err := c.GetConfiguration(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_configuration.getConfiguration", "configuration", fmt.Sprintf("%+v", configuration))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_configuration.getConfiguration", "request_error", err)
		return nil, err
	}

	return configuration, nil
}
