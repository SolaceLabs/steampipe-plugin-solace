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
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventBrokerServices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventBrokerServices", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventBrokerServiceListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventBrokerServices, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service.listEventBrokerServices", "eventBrokerServices", eventBrokerServices)
		if err != nil {
			plugin.Logger(ctx).Error("solace_event_broker_service.listEventBrokerServices", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service.getEventBrokerService - ID", id)

	eventBrokerService, err := c.GetEventBrokerService(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event_broker_service.getEventBrokerService", "eventBrokerService", eventBrokerService)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event_broker_service.getEventBrokerService", "request_error", err)
		return nil, err
	}

	return eventBrokerService, nil
}
