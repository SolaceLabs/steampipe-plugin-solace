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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Application Domain Id."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Application Domain Name."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description."},
			{Name: "unique_topic_address_enforcement_enabled", Type: proto.ColumnType_BOOL, Description: "Unique Topic Enforcement enabled?."},
			{Name: "topic_domain_enforcement_enabled", Type: proto.ColumnType_BOOL, Description: "Topic Domain Enforcement Enabled?."},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: "Stats."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom Attributes."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Created By."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Created Time."},
			{Name: "changed_by", Type: proto.ColumnType_STRING, Description: "Changed By."},
			{Name: "updated_time", Type: proto.ColumnType_TIMESTAMP, Description: "Updated Time."},
		},
	}
}

func listApplicationDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listApplicationDomains", ctx, d, h)
	plugin.Logger(ctx).Debug("DEBUGGING solace_application_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewApplicationDomainListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		applicationDomains, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Debug("DEBUGGING solace_application_domain.listApplicationDomains", "applicationDomains", applicationDomains)
		if err != nil {
			plugin.Logger(ctx).Error("solace_application_domain.listApplicationDomains", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Debug("RECORDS FETCHED - ", count)
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
	plugin.Logger(ctx).Debug("DEBUGGING solace_application_domain.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Debug("DEBUGGING solace_application_domain.getApplicationDomain - ID", id)

	applicationDomain, err := c.GetApplicationDomain(id)
	plugin.Logger(ctx).Debug("DEBUGGING solace_application_domain.getApplicationDomain", "applicationDomain", applicationDomain)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_application_domain.getApplicationDomain", "request_error", err)
		return nil, err
	}

	return applicationDomain, nil
}
