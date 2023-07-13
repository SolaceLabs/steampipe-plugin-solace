package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventApiProductVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_eventapi_product_version",
		Description: "Get a list of eventApiProductVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventApiProductVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventApiProductVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventApiProductId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "summary", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventApiVersionIds", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "eventApiProductRegistrations", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "solaceClassOfServicePolicy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "plans", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "solaceMessagingServices", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "topicFilters", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "filters", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "approvalType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "publishState", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "publishedTime", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventApiProductVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiProductVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProductVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiProductVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiProductVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProductVersion.listEventApiProductVersions", "eventApiProductVersions", eventApiProductVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiProductVersion.listEventApiProductVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventApiProductVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventApiProductVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProductVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")

	eventApiProductVersion, err := c.GetEventApiProductVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProductVersion.getEventApiProductVersion", "eventApiProductVersion", fmt.Sprintf("%+v", eventApiProductVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiProductVersion.getEventApiProductVersion", "request_error", err)
		return nil, err
	}

	return eventApiProductVersion, nil
}
