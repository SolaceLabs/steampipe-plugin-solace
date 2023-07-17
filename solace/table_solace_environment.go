package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_environment",
		Description: "Get a list of environments that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEnvironments,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEnvironment,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description"},
			{Name: "revision", Type: proto.ColumnType_INT, Description: "Revision"},
			{Name: "number_of_event_meshes", Type: proto.ColumnType_INT, Description: "Number of Event Meshes"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEnvironments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEnvironments", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_environment.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEnvironmentListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		environments, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_environment.listEnvironments", "environments", environments)
		if err != nil {
			plugin.Logger(ctx).Error("solace_environment.listEnvironments", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range environments {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEnvironment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Debug("DEBUGGING solace_environment.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_environment.getEnvironment - ID", id)

	environment, err := c.GetEnvironment(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_environment.getEnvironment", "environment", fmt.Sprintf("%+v", environment))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_environment.getEnvironment", "request_error", err)
		return nil, err
	}

	return environment, nil
}
