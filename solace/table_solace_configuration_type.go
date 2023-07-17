package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableConfigurationType(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_configuration_type",
		Description: "Get a list of configuration types that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listConfigurationTypes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getConfigurationType,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "broker_type", Type: proto.ColumnType_STRING, Description: "Broker Type"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "associated_entity_types", Type: proto.ColumnType_STRING, Description: "Associated Entity Types"},
			{Name: "value_schema", Type: proto.ColumnType_JSON, Description: "Value Schema"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listConfigurationTypes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listConfigurationTypes", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_configurationType.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewConfigurationTypeListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		configurationTypes, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_configurationType.listConfigurationTypes", "configurationTypes", configurationTypes)
		if err != nil {
			plugin.Logger(ctx).Error("solace_configurationType.listConfigurationTypes", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range configurationTypes {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getConfigurationType(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_configurationType.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_configurationType.getConfigurationType - ID", id)

	configurationType, err := c.GetConfigurationType(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_configurationType.getConfigurationType", "configurationType", fmt.Sprintf("%+v", configurationType))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_configurationType.getConfigurationType", "request_error", err)
		return nil, err
	}

	return configurationType, nil
}
