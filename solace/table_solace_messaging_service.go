package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableMessagingService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_messaging_service",
		Description: "Get a list of messagingServices that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listMessagingServices,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMessagingService,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "event_mesh_id", Type: proto.ColumnType_STRING, Description: "Event Mesh Id"},
			{Name: "context_id", Type: proto.ColumnType_STRING, Description: "Context Id"},
			{Name: "runtime_agent_id", Type: proto.ColumnType_STRING, Description: "Runtime Agent Id"},
			{Name: "solace_cloud_messaging_service_id", Type: proto.ColumnType_STRING, Description: "Solace Cloud Messaging Service Id"},
			{Name: "messaging_service_type", Type: proto.ColumnType_STRING, Description: "Messaging Service Type"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "messaging_service_connections", Type: proto.ColumnType_JSON, Description: "Messaging Service Connections"},
			{Name: "event_management_agent_id", Type: proto.ColumnType_STRING, Description: "Event Management Agent Id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listMessagingServices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listMessagingServices", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_messagingService.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewMessagingServiceListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		messagingServices, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_messagingService.listMessagingServices", "messagingServices", messagingServices)
		if err != nil {
			plugin.Logger(ctx).Error("solace_messagingService.listMessagingServices", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range messagingServices {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getMessagingService(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_messagingService.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_messagingService.getMessagingService - ID", id)

	messagingService, err := c.GetMessagingService(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_messagingService.getMessagingService", "messagingService", fmt.Sprintf("%+v", messagingService))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_messagingService.getMessagingService", "request_error", err)
		return nil, err
	}

	return messagingService, nil
}
