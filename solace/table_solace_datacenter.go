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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name."},
			{Name: "datacenter_type", Type: proto.ColumnType_STRING, Description: "Datacenter type."},
			{Name: "provider", Type: proto.ColumnType_STRING, Description: "Provider."},
			{Name: "oper_state", Type: proto.ColumnType_STRING, Description: "Operational State."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "updated_by", Type: proto.ColumnType_STRING, Description: "Updated By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Updated Time."},
			{Name: "available", Type: proto.ColumnType_STRING, Description: "Available?."},
			{Name: "supported_service_classes", Type: proto.ColumnType_STRING, Description: "Supported Service Classes."},
			{Name: "cloud_agent_version", Type: proto.ColumnType_STRING, Description: "Cloud Agent Version."},
			{Name: "k8s_service_type", Type: proto.ColumnType_STRING, Description: "K8s Service Type."},
			{Name: "num_supported_private_endpoints", Type: proto.ColumnType_STRING, Description: "Number of Supported Private Endpoints."},
			{Name: "num_supported_public_endpoints", Type: proto.ColumnType_STRING, Description: "Number of Supported Public Endpoints."},
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Organization Id."},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Location."},
			{Name: "spool_scale_up_capability_info", Type: proto.ColumnType_JSON, Description: "Spool Scale Up Capability Info."},
		},
	}
}

func listDatacenters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listDatacenters", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_datacenter.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewDatacenterListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		applications, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_datacenter.listDatacenters", "applications", applications)
		if err != nil {
			plugin.Logger(ctx).Error("solace_datacenter.listDatacenters", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_datacenter.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_datacenter.getDatacenter - ID", id)

	datacenter, err := c.GetDatacenter(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_datacenter.getDatacenter", "datacenter", datacenter)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_datacenter.getDatacenter", "request_error", err)
		return nil, err
	}

	return datacenter, nil
}
