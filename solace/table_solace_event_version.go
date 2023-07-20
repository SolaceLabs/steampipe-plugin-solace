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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "event_id", Type: proto.ColumnType_STRING, Description: "Event Id."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display Name."},
			{Name: "declared_producing_application_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredProducingApplicationVersionIdsAsString"), Description: "Declared Producing Application Version Ids."},
			{Name: "declared_consuming_application_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredConsumingApplicationVersionIdsAsString"), Description: "Declared Consuming Application Version Ids."},
			{Name: "producing_event_api_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("ProducingEventApiVersionIdsAsString"), Description: "Producing Event API Version Ids."},
			{Name: "consuming_event_api_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("ConsumingEventApiVersionIdsAsString"), Description: "Consuming Event API Version Ids."},
			{Name: "attracting_application_version_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("AttractingApplicationVersionIdsAsString"), Description: "Attracting Application Version Ids."},
			{Name: "delivery_descriptor", Type: proto.ColumnType_JSON, Description: "Delivery Descriptor."},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "State Id."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "messaging_service_ids", Type: proto.ColumnType_STRING, Description: "Message Service Ids."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listEventVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventVersion.listEventVersions", "eventVersions", eventVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventVersion.listEventVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventVersion.getEventVersion - ID", id)

	eventVersion, err := c.GetEventVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventVersion.getEventVersion", "eventVersion", fmt.Sprintf("%+v", eventVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventVersion.getEventVersion", "request_error", err)
		return nil, err
	}

	return eventVersion, nil
}
