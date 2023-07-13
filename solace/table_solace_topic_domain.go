package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTopicDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_topic_domain",
		Description: "Get a list of topic domains that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listTopicDomains,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTopicDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "applicationDomainId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "brokerType", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "addressLevels", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

func listTopicDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listTopicDomains", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_topic_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewTopicDomainListPaginator()
	pagesLeft := true
	count := 0
	for pagesLeft {
		topicDomains, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_topic_domain.listTopicDomains", "topicDomains", topicDomains)
		if err != nil {
			plugin.Logger(ctx).Error("solace_topic_domain.listTopicDomains", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range topicDomains {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getTopicDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_topic_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_topic_domain.getTopicDomain - ID", id)

	topicDomain, err := c.GetTopicDomain(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_topic_domain.getTopicDomain", "topicDomain", fmt.Sprintf("%+v", topicDomain))
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_topic_domain.getTopicDomain", "request_error", err)
		return nil, err
	}

	return topicDomain, nil
}
