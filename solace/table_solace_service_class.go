package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableServiceClass(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_service_class",
		Description: "Get a list of service classes that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listServiceClasses,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getServiceClass,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "vpnConnections", Type: proto.ColumnType_INT, Description: ""},
			{Name: "brokerScalingTier", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "vpnMaxSpoolSize", Type: proto.ColumnType_INT, Description: ""},
			{Name: "maxNumberVpns", Type: proto.ColumnType_INT, Description: ""},
			{Name: "highAvailabilityCapable", Type: proto.ColumnType_BOOL, Description: ""},
		},
	}
}

func listServiceClasses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listServiceClasses", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_service_class.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewServiceClassListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		serviceClasses, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_service_class.listServiceClasses", "serviceClasses", serviceClasses)
		if err != nil {
			plugin.Logger(ctx).Error("solace_service_class.listServiceClasses", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range serviceClasses {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getServiceClass(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_service_class.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_service_class.getServiceClass - ID", id)

	serviceClass, err := c.GetServiceClass(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_service_class.getServiceClass", "serviceClass", fmt.Sprintf("%+v", serviceClass))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_service_class.getServiceClass", "request_error", err)
		return nil, err
	}

	return serviceClass, nil
}
