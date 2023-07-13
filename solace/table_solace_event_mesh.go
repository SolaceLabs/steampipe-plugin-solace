package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventMesh(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_event_mesh",
		Description: "Get a list of event meshes that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventMeshes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventMesh,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventMeshes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventMeshes", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventMesh.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventMeshListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventMeshes, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventMesh.listEventMeshes", "eventMeshes", eventMeshes)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventMesh.listEventMeshes", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventMeshes {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventMesh(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventMesh.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventMesh.getEventMesh - ID", id)

	eventMesh, err := c.GetEventMesh(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventMesh.getEventMesh", "eventMesh", fmt.Sprintf("%+v", eventMesh))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventMesh.getEventMesh", "request_error", err)
		return nil, err
	}

	return eventMesh, nil
}
