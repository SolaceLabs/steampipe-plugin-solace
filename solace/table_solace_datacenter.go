package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDatacenter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_datacenter",
		Description: "Get a list of datacenters that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listDatacenters,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDatacenter,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "datacenterType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "provider", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "operState", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "available", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "supportedServiceClasses", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "cloudAgentVersion", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "k8sServiceType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numSupportedPrivateEndpoints", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numSupportedPublicEndpoints", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "organizationId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "location", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "spoolScaleUpCapabilityInfo", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func listDatacenters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listDatacenters", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_datacenter.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewDatacenterListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		applications, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_datacenter.listDatacenters", "applications", applications)
		if err != nil {
			plugin.Logger(ctx).Error("solace_datacenter.listDatacenters", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range applications {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getDatacenter(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_datacenter.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_datacenter.getDatacenter - ID", id)

	datacenter, err := c.GetDatacenter(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_datacenter.getDatacenter", "datacenter", datacenter)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_datacenter.getDatacenter", "request_error", err)
		return nil, err
	}

	return datacenter, nil
}
