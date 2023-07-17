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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "event_api_id", Type: proto.ColumnType_STRING, Description: "Event API Id"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description"},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display Name"},
			{Name: "produced_event_version_ids", Type: proto.ColumnType_STRING, Description: "Produced Event Version Ids"},
			{Name: "consumed_event_version_ids", Type: proto.ColumnType_STRING, Description: "Consumed Event Version Ids"},
			{Name: "declared_event_api_product_version_ids", Type: proto.ColumnType_STRING, Description: "Declared Event API Product Version Ids"},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes"},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "State Id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEventApiVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiVersion.listEventApiVersions", "eventApiVersions", eventApiVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiVersion.listEventApiVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiVersion.getEventApiVersion - ID", id)

	eventApiVersion, err := c.GetEventApiVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiVersion.getEventApiVersion", "eventApiVersion", fmt.Sprintf("%+v", eventApiVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiVersion.getEventApiVersion", "request_error", err)
		return nil, err
	}

	return eventApiVersion, nil
}
