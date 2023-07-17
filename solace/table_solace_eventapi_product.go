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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "application_domain_id", Type: proto.ColumnType_STRING, Description: "Application Domain Id"},
			{Name: "shared", Type: proto.ColumnType_STRING, Description: "Shared?"},
			{Name: "number_of_versions", Type: proto.ColumnType_STRING, Description: "Number of Versions"},
			{Name: "broker_type", Type: proto.ColumnType_STRING, Description: "Broker Type"},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Object type"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By"},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time"},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Modified By"},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Modified Time"},
		},
	}
}

func listEventApiProducts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listEventApiProducts", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProduct.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewEventApiProductListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		eventApiProducts, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProduct.listEventApiProducts", "eventApiProducts", eventApiProducts)
		if err != nil {
			plugin.Logger(ctx).Error("solace_eventApiProduct.listEventApiProducts", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProduct.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProduct.getEventApiProduct - ID", id)

	eventApiProduct, err := c.GetEventApiProduct(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_eventApiProduct.getEventApiProduct", "eventApiProduct", fmt.Sprintf("%+v", eventApiProduct))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_eventApiProduct.getEventApiProduct", "request_error", err)
		return nil, err
	}

	return eventApiProduct, nil
}
