package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventBrokerServiceVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_broker_service_version",
		Description: "Get a event broker service version that match the given parameters",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventBrokerServiceVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "supportedServiceClasses", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "capabilities", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func getEventBrokerServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_version.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_version.getEventBrokerServiceVersion for datacenterId - ", id)

	eventBrokerServiceVersions, err := c.GetEventBrokerServiceVersions(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_version.getEventBrokerServiceVersions", "eventBrokerServiceVersions", eventBrokerServiceVersions)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event_broker_service_version.getEventBrokerServiceVersions", "request_error", err)
		return nil, err
	}

	for _, i := range eventBrokerServiceVersions {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
