package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventManagementAgent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_management_agent",
		Description: "Get a list of event management agents that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventManagementAgents,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventManagementAgent,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Region"},
			{Name: "client_username", Type: proto.ColumnType_STRING, Description: "Client Username"},
			{Name: "client_password", Type: proto.ColumnType_STRING, Description: "Client Password"},
			{Name: "referenced_by_messaging_service_ids", Type: proto.ColumnType_STRING, Description: "Referenced by Messaging Service Ids"},
			{Name: "org_id", Type: proto.ColumnType_STRING, Description: "Organization Id"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status"},
			{Name: "last_connected_time", Type: proto.ColumnType_STRING, Description: "Last Connected Time"},
			{Name: "event_management_agent_region_id", Type: proto.ColumnType_STRING, Description: "Event Management Agent Region Id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEventManagementAgents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventManagementAgents", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventManagementAgent.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventManagementAgentListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventManagementAgents, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventManagementAgent.listEventManagementAgents", "eventManagementAgents", eventManagementAgents)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventManagementAgent.listEventManagementAgents", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventManagementAgents {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventManagementAgent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventManagementAgent.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventManagementAgent.getEventManagementAgent - ID", id)

	eventManagementAgent, err := c.GetEventManagementAgent(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventManagementAgent.getEventManagementAgent", "eventManagementAgent", fmt.Sprintf("%+v", eventManagementAgent))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventManagementAgent.getEventManagementAgent", "request_error", err)
		return nil, err
	}

	return eventManagementAgent, nil
}
