package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableApplicationVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_application_version",
		Description: "Get a list of applicationVersions that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listApplicationVersions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getApplicationVersion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Application Version Id."},
			{Name: "application_id", Type: proto.ColumnType_STRING, Description: "Application Id."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version Number."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Version Display Name."},
			{Name: "declared_produced_event_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredProducedEventVersionIdsAsString"), Description: "Event Version IDs produced by the Application."},
			{Name: "declared_consumed_event_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredConsumedEventVersionIdsAsString"), Description: "Event Version IDs consumed by the Application."},
			{Name: "declared_event_api_product_version_ids", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredEventApiProductVersionIdsAsString"), Description: "Event API Product Version Ids."},
			{Name: "state_id", Type: proto.ColumnType_STRING, Description: "Application State."},
			{Name: "consumers", Type: proto.ColumnType_JSON, Description: "Application Consumers."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "messaging_service_ids", Type: proto.ColumnType_STRING, Description: "Messaging Service Ids."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time."},
		},
	}
}

func listApplicationVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listApplicationVersions", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_applicationVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewApplicationVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		applicationVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_applicationVersion.listApplicationVersions", "applicationVersions", applicationVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_applicationVersion.listApplicationVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range applicationVersions {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getApplicationVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_applicationVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_applicationVersion.getApplicationVersion - ID", id)

	applicationVersion, err := c.GetApplicationVersion(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_applicationVersion.getApplicationVersion", "applicationVersion", fmt.Sprintf("%+v", applicationVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_applicationVersion.getApplicationVersion", "request_error", err)
		return nil, err
	}

	return applicationVersion, nil
}
