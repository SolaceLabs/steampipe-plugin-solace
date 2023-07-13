package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventApiVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_eventapi_version",
		Description: "Get a list of eventApiVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventApiVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventApiVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventApiId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "producedEventVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "consumedEventVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "declaredEventApiProductVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventApiVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiVersion.listEventApiVersions", "eventApiVersions", eventApiVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiVersion.listEventApiVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventApiVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventApiVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiVersion.getEventApiVersion - ID", id)

	eventApiVersion, err := c.GetEventApiVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiVersion.getEventApiVersion", "eventApiVersion", fmt.Sprintf("%+v", eventApiVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiVersion.getEventApiVersion", "request_error", err)
		return nil, err
	}

	return eventApiVersion, nil
}
