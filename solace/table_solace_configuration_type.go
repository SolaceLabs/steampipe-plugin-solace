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
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "associatedEntityTypes", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "valueSchema", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listConfigurationTypes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listConfigurationTypes", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_configurationType.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewConfigurationTypeListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		configurationTypes, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_configurationType.listConfigurationTypes", "configurationTypes", configurationTypes)
		if err != nil {
			plugin.Logger(ctx).Error("solace_configurationType.listConfigurationTypes", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Trace("DEBUGGING solace_configurationType.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_configurationType.getConfigurationType - ID", id)

	configurationType, err := c.GetConfigurationType(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_configurationType.getConfigurationType", "configurationType", fmt.Sprintf("%+v", configurationType))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_configurationType.getConfigurationType", "request_error", err)
		return nil, err
	}

	return configurationType, nil
}
