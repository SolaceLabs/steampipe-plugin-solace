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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "event_api_product_id", Type: proto.ColumnType_STRING, Description: "Event API Product Id."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version."},
			{Name: "summary", Type: proto.ColumnType_STRING, Description: "Summary."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display Name."},
			{Name: "event_api_version_ids", Type: proto.ColumnType_STRING, Description: "Event API Version Ids."},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "State Id."},
			{Name: "event_api_product_registrations", Type: proto.ColumnType_STRING, Description: "Event API Product Registrations."},
			{Name: "solace_class_of_service_policy", Type: proto.ColumnType_STRING, Description: "Solace Class of Service Policy."},
			{Name: "plans", Type: proto.ColumnType_STRING, Description: "Plans."},
			{Name: "solace_messaging_services", Type: proto.ColumnType_STRING, Description: "Solace Messaging Services."},
			{Name: "topic_filters", Type: proto.ColumnType_STRING, Description: "Topic Filters."},
			{Name: "filters", Type: proto.ColumnType_STRING, Description: "Filters."},
			{Name: "approval_type", Type: proto.ColumnType_STRING, Description: "Approval Type."},
			{Name: "publish_state", Type: proto.ColumnType_STRING, Description: "Publish State."},
			{Name: "published_time", Type: proto.ColumnType_STRING, Description: "Published Time."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listEventApiProductVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiProductVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProductVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiProductVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiProductVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProductVersion.listEventApiProductVersions", "eventApiProductVersions", eventApiProductVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiProductVersion.listEventApiProductVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProductVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")

	eventApiProductVersion, err := c.GetEventApiProductVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProductVersion.getEventApiProductVersion", "eventApiProductVersion", fmt.Sprintf("%+v", eventApiProductVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiProductVersion.getEventApiProductVersion", "request_error", err)
		return nil, err
	}

	return eventApiProductVersion, nil
}
