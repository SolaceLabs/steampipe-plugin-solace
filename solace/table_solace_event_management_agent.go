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
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "region", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "clientUsername", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "clientPassword", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "referencedByMessagingServiceIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "orgId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "lastConnectedTime", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventManagementAgentRegionId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventManagementAgents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventManagementAgents", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventManagementAgent.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventManagementAgentListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventManagementAgents, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventManagementAgent.listEventManagementAgents", "eventManagementAgents", eventManagementAgents)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventManagementAgent.listEventManagementAgents", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventManagementAgent.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventManagementAgent.getEventManagementAgent - ID", id)

	eventManagementAgent, err := c.GetEventManagementAgent(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventManagementAgent.getEventManagementAgent", "eventManagementAgent", fmt.Sprintf("%+v", eventManagementAgent))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventManagementAgent.getEventManagementAgent", "request_error", err)
		return nil, err
	}

	return eventManagementAgent, nil
}
