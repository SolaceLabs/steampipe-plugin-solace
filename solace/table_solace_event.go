package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event",
		Description: "Get a list of events that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEvents,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEvent,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationDomainId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numberOfVersions", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEvents", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		events, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_event.listEvents", "events", events)
		if err != nil {
			plugin.Logger(ctx).Error("solace_event.listEvents", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range events {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEvent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_event.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")

	event, err := c.GetEvent(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_event.getEvent", "event", fmt.Sprintf("%+v", event))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_event.getEvent", "request_error", err)
		return nil, err
	}

	return event, nil
}
