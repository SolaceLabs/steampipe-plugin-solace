package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableConsumer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_consumer",
		Description: "Get a list of consumers that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listConsumers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getConsumer,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "consumerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationVersionId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "subscriptions", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listConsumers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listConsumers", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_consumer.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewConsumerListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		consumers, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_consumer.listConsumers", "consumers", consumers)
		if err != nil {
			plugin.Logger(ctx).Error("solace_consumer.listConsumers", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range consumers {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getConsumer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_consumer.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_consumer.getConsumer - ID", id)

	consumer, err := c.GetConsumer(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_consumer.getConsumer", "consumer", fmt.Sprintf("%+v", consumer))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_consumer.getConsumer", "request_error", err)
		return nil, err
	}

	return consumer, nil
}
