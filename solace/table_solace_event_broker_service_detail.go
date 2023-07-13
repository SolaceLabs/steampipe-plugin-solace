package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventBrokerServiceDetail(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_broker_service_detail",
		Description: "Get a event broker service detail that match the given parameters",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventBrokerServiceDetail,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ownedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "infrastructureId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "datacenterId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "serviceClassId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventMeshId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ongoingOperationIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "adminState", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "creationState", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "locked", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "defaultManagementHostname", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "serviceConnectionEndpoints", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "broker", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func getEventBrokerServiceDetail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_detail.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail - ID", id)

	eventBrokerServiceDetail, err := c.GetEventBrokerServiceDetail(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail", "eventBrokerServiceDetail", eventBrokerServiceDetail)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail", "request_error", err)
		return nil, err
	}

	return eventBrokerServiceDetail, nil
}
