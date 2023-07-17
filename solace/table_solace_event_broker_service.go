package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventBrokerService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_broker_service",
		Description: "Get a list of event broker services that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventBrokerServices,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventBrokerService,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "owned_by", Type: proto.ColumnType_STRING, Description: "Owned By"},
			{Name: "infrastructure_id", Type: proto.ColumnType_STRING, Description: "Infrastructure Id"},
			{Name: "datacenter_id", Type: proto.ColumnType_STRING, Description: "Datacenter Id"},
			{Name: "service_class_id", Type: proto.ColumnType_STRING, Description: "Service Class Id"},
			{Name: "event_mesh_id", Type: proto.ColumnType_STRING, Description: "Event Mesh Id"},
			{Name: "ongoing_operation_ids", Type: proto.ColumnType_STRING, Description: "Ongoing Operation Ids"},
			{Name: "admin_state", Type: proto.ColumnType_STRING, Description: "Admin State"},
			{Name: "creation_state", Type: proto.ColumnType_STRING, Description: "Creation State"},
			{Name: "locked", Type: proto.ColumnType_STRING, Description: "Locked?"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEventBrokerServices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventBrokerServices", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventBrokerServiceListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventBrokerServices, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service.listEventBrokerServices", "eventBrokerServices", eventBrokerServices)
		if err != nil {
			plugin.Logger(ctx).Error("solace_event_broker_service.listEventBrokerServices", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventBrokerServices {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventBrokerService(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service.getEventBrokerService - ID", id)

	eventBrokerService, err := c.GetEventBrokerService(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_event_broker_service.getEventBrokerService", "eventBrokerService", eventBrokerService)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event_broker_service.getEventBrokerService", "request_error", err)
		return nil, err
	}

	return eventBrokerService, nil
}
