package solace

import (
	"context"
	"fmt"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDomainDependency(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "solace_domain_dependency",
		Description: "Get a event portal domain dependencies that match the given parameters",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    listDomainDependencies,
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("id"),
		// 	Hydrate:    getDomainDependency,
		// },
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "resourceId", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "resource", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "application", Type: proto.ColumnType_JSON, Description: ""},
			// {Name: "event", Type: proto.ColumnType_JSON, Description: ""},
			// {Name: "schema", Type: proto.ColumnType_JSON, Description: ""},
			// {Name: "enum", Type: proto.ColumnType_JSON, Description: ""},
			// {Name: "eventapi", Type: proto.ColumnType_JSON, Description: ""},
			// {Name: "eventapiproduct", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func listDomainDependencies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	LogQueryContext("listDomainDependencies", ctx, d, h)
	plugin.Logger(ctx).Trace("DEBUGGING solace_domain_dependency.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))

	client, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}

	tlp := client.NewDomainDependencyListPaginator(nil)
	pagesLeft := true
	count := 0
	for pagesLeft {
		dependencies, meta, err := tlp.NextPage()
		plugin.Logger(ctx).Trace("DEBUGGING solace_domain_dependency.listDomainDependencies", "dependencies", dependencies)
		if err != nil {
			plugin.Logger(ctx).Error("solace_domain_dependency.listDomainDependencies", "request_error", err)
			pagesLeft = false
			// return nil, err
		} else {
			count += meta.Pagination.Count
			plugin.Logger(ctx).Trace("RECORDS FETCHED - ", count)
		}

		log.Println("DOMAIN HIERARCHY", fmt.Sprintf("%+v", dependencies))

		// stream results
		for _, i := range dependencies {
			d.StreamListItem(ctx, i)
			log.Println("HIERARCHY RECORD", fmt.Sprintf("%v", i))

			if d.RowsRemaining(ctx) <= 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getDomainDependency(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("DEBUGGING solace_domain_dependency.QueryData", "QueryData", fmt.Sprintf("%+v", d.FetchType))
	c, err := NewSolaceClient(d.Connection)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQualString("id")
	plugin.Logger(ctx).Trace("DEBUGGING solace_domain_dependency.getDomainDependency - ID", id)

	domainDependencies, err := c.GetDomainDependency(id)
	plugin.Logger(ctx).Trace("DEBUGGING solace_domain_dependency.getDomainDependency", "applicationDomain", domainDependencies)
	if err != nil {
		plugin.Logger(ctx).Error("DEBUGGING solace_domain_dependency.getDomainDependency", "request_error", err)
		return nil, err
	}

	// stream results
	// for _, i := range domainDependencies {
	// 	d.StreamListItem(ctx, i)

	// 	if d.RowsRemaining(ctx) <= 0 {
	// 		return nil, nil
	// 	}
	// }

	return domainDependencies, nil
}
