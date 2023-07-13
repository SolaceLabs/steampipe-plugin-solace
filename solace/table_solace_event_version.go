package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableEventVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_version",
		Description: "Get a list of eventVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "declaredProducingApplicationVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredProducingApplicationVersionIdsAsString"), Description: ""},
			{Name: "declaredConsumingApplicationVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredConsumingApplicationVersionIdsAsString"), Description: ""},
			{Name: "producingEventApiVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("ProducingEventApiVersionIdsAsString"), Description: ""},
			{Name: "consumingEventApiVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("ConsumingEventApiVersionIdsAsString"), Description: ""},
			{Name: "attractingApplicationVersionIds", Type: proto.ColumnType_JSON, Transform: transform.FromField("AttractingApplicationVersionIdsAsString"), Description: ""},
			{Name: "deliveryDescriptor", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "messagingServiceIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventVersion.listEventVersions", "eventVersions", eventVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventVersion.listEventVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventVersion.getEventVersion - ID", id)

	eventVersion, err := c.GetEventVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventVersion.getEventVersion", "eventVersion", fmt.Sprintf("%+v", eventVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventVersion.getEventVersion", "request_error", err)
		return nil, err
	}

	return eventVersion, nil
}
