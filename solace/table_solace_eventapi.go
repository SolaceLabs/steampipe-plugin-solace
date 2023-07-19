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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name."},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: "Shared?."},
			{Name: "application_domain_id", Type: proto.ColumnType_STRING, Description: "Application Domain Id."},
			{Name: "number_of_versions", Type: proto.ColumnType_STRING, Description: "Number of Versions."},
			{Name: "broker_type", Type: proto.ColumnType_STRING, Description: "Broker Type."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listEventApis(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApis", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApi.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApis, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventApi.listEventApis", "eventApis", eventApis)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApi.listEventApis", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApi.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApi.getEventApi - ID", id)

	eventApi, err := c.GetEventApi(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApi.getEventApi", "eventApi", fmt.Sprintf("%+v", eventApi))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApi.getEventApi", "request_error", err)
		return nil, err
	}

	return eventApi, nil
}
