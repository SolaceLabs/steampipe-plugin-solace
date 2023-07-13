package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventApi(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_eventapi",
		Description: "Get a list of eventApis that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventApis,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventApi,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationDomainId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numberOfVersions", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventApis(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApis", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApi.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApis, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventApi.listEventApis", "eventApis", eventApis)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApi.listEventApis", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventApis {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventApi(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApi.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApi.getEventApi - ID", id)

	eventApi, err := c.GetEventApi(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApi.getEventApi", "eventApi", fmt.Sprintf("%+v", eventApi))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApi.getEventApi", "request_error", err)
		return nil, err
	}

	return eventApi, nil
}
