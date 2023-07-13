package solace

import (
	"context"
	"fmt"
	"strings"

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
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "declaredProducedEventVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredProducedEventVersionIdsAsString"), Description: ""},
			{Name: "declaredConsumedEventVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredConsumedEventVersionIdsAsString"), Description: ""},
			{Name: "declaredEventApiProductVersionIds", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeclaredEventApiProductVersionIdsAsString"), Description: ""},
			{Name: "stateId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "consumers", Type: proto.ColumnType_JSON, Description: ""},
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

func buildCSV(ctx context.Context, d *transform.TransformData) (string, error) {
	fields := d.Value.([]string)
	plugin.Logger(ctx).Trace("DEBUGGING TRANSFORMP", fields)
	if fields != nil {
		return strings.Join(fields, ", "), nil
	}
	plugin.Logger(ctx).Trace("DEBUGGING TRANSFORMP EMPTY", fields)
	return "", nil
}

func listApplicationVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listApplicationVersions", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_applicationVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewApplicationVersionListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		applicationVersions, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_applicationVersion.listApplicationVersions", "applicationVersions", applicationVersions)
		if err != nil {
			plugin.Logger(ctx).Error("solace_applicationVersion.listApplicationVersions", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Trace("DEBUGGING solace_applicationVersion.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_applicationVersion.getApplicationVersion - ID", id)

	applicationVersion, err := c.GetApplicationVersion(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_applicationVersion.getApplicationVersion", "applicationVersion", fmt.Sprintf("%+v", applicationVersion))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_applicationVersion.getApplicationVersion", "request_error", err)
		return nil, err
	}

	return applicationVersion, nil
}
