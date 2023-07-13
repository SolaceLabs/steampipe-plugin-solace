package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLifecycleState(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_lifecycle_state",
		Description: "Get a list of lifecycle states that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listLifecycleStates,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "stateOrder", Type: proto.ColumnType_INT, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listLifecycleStates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listLifecycleStates", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_lifecycle_state.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewLifecycleStateListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		lifecycleStates, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_lifecycle_state.listLifecycleStates", "lifecycleStates", lifecycleStates)
		if err != nil {
			plugin.Logger(ctx).Error("solace_lifecycle_state.listLifecycleStates", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range lifecycleStates {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
