package solace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableApplicationDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_application_domain",
		Description: "Get a list of application domains that match the given parameters",
		List: &plugin.ListConfig{
			Hydrate: listApplicationDomains,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getApplicationDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "uniqueTopicAddressEnforcementEnabled", Type: proto.ColumnType_BOOL, Description: "Forces all topic addresses within the application domain to be unique."},
			{Name: "topicDomainEnforcementEnabled", Type: proto.ColumnType_BOOL, Description: "Forces all topic addresses within the application domain to be prefixed with one of the application domain’s configured topic domains."},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "customAttributes", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "createdBy", Type: proto.ColumnType_STRING, Description: "Event broker service capabilities."},
			{Name: "createdTime", Type: proto.ColumnType_TIMESTAMP, Description: "Event broker service version."},
			{Name: "changedBy", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updatedTime", Type: proto.ColumnType_TIMESTAMP, Description: "Supported service classes."},
		},
	}
}

func listApplicationDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listApplicationDomains", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_application_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewApplicationDomainListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		applicationDomains, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_application_domain.listApplicationDomains", "applicationDomains", applicationDomains)
		if err != nil {
			plugin.Logger(ctx).Error("solace_application_domain.listApplicationDomains", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		// stream results
		for _, i := range applicationDomains {
			d.StreamListItem(ctx, i)

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getApplicationDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_application_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_application_domain.getApplicationDomain - ID", id)

	applicationDomain, err := c.GetApplicationDomain(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_application_domain.getApplicationDomain", "applicationDomain", applicationDomain)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_application_domain.getApplicationDomain", "request_error", err)
		return nil, err
	}

	return applicationDomain, nil
}
