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
		Description: "Get event broker service detail that match the given parameters",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventBrokerServiceDetail,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name."},
			{Name: "owned_by", Type: proto.ColumnType_STRING, Description: "Owned By."},
			{Name: "infrastructure_id", Type: proto.ColumnType_STRING, Description: "Infrastructure Id."},
			{Name: "datacenter_id", Type: proto.ColumnType_STRING, Description: "Datacenter Id."},
			{Name: "service_class_id", Type: proto.ColumnType_STRING, Description: "Service Class Id."},
			{Name: "event_mesh_id", Type: proto.ColumnType_STRING, Description: "Event Mesh Id."},
			{Name: "ongoing_operation_ids", Type: proto.ColumnType_STRING, Description: "Ongoing Operation Ids."},
			{Name: "admin_state", Type: proto.ColumnType_STRING, Description: "Admin State."},
			{Name: "creation_state", Type: proto.ColumnType_STRING, Description: "Creation State."},
			{Name: "locked", Type: proto.ColumnType_STRING, Description: "Locked?."},
			{Name: "default_management_hostname", Type: proto.ColumnType_STRING, Description: "Default Management Hostname."},
			{Name: "service_connection_endpoints", Type: proto.ColumnType_JSON, Description: "Service Connection Endpoints."},
			{Name: "broker", Type: proto.ColumnType_JSON, Description: "Broker."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func getEventBrokerServiceDetail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service_detail.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail - ID", id)

	eventBrokerServiceDetail, err := c.GetEventBrokerServiceDetail(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail", "eventBrokerServiceDetail", eventBrokerServiceDetail)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event_broker_service_detail.getEventBrokerServiceDetail", "request_error", err)
		return nil, err
	}

	return eventBrokerServiceDetail, nil
}
