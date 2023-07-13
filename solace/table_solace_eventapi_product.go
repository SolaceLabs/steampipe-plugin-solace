package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEventApiProduct(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_eventapi_product",
		Description: "Get a list of eventApiProducts that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listEventApiProducts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEventApiProduct,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationDomainId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "numberOfVersions", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listEventApiProducts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiProducts", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProduct.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiProductListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiProducts, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProduct.listEventApiProducts", "eventApiProducts", eventApiProducts)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiProduct.listEventApiProducts", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range eventApiProducts {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getEventApiProduct(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProduct.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProduct.getEventApiProduct - ID", id)

	eventApiProduct, err := c.GetEventApiProduct(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_eventApiProduct.getEventApiProduct", "eventApiProduct", fmt.Sprintf("%+v", eventApiProduct))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiProduct.getEventApiProduct", "request_error", err)
		return nil, err
	}

	return eventApiProduct, nil
}
