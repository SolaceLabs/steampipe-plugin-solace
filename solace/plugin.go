package solace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-solace"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"solace_application_domain":          tableApplicationDomain(ctx),
			"solace_application":                 tableApplication(ctx),
			"solace_application_version":         tableApplicationVersion(ctx),
			"solace_event":                       tableEvent(ctx),
			"solace_event_version":               tableEventVersion(ctx),
			"solace_schema":                      tableSchema(ctx),
			"solace_schema_version":              tableSchemaVersion(ctx),
			"solace_enum":                        tableEnum(ctx),
			"solace_enum_version":                tableEnumVersion(ctx),
			"solace_eventapi":                    tableEventApi(ctx),
			"solace_eventapi_version":            tableEventApiVersion(ctx),
			"solace_eventapi_product":            tableEventApiProduct(ctx),
			"solace_eventapi_product_version":    tableEventApiProductVersion(ctx),
			"solace_event_broker_service":        tableEventBrokerService(ctx),
			"solace_event_broker_service_detail": tableEventBrokerServiceDetail(ctx),
			"solace_topic_domain":                tableTopicDomain(ctx),
			"solace_lifecycle_state":             tableLifecycleState(ctx),
			"solace_consumer":                    tableConsumer(ctx),
			"solace_datacenter":                  tableDatacenter(ctx),
			"solace_service_class":               tableServiceClass(ctx),
			"solace_event_mesh":                  tableEventMesh(ctx),
			"solace_environment":                 tableEnvironment(ctx),
			"solace_configuration":               tableConfiguration(ctx),
			"solace_messaging_service":           tableMessagingService(ctx),
			"solace_configuration_type":          tableConfigurationType(ctx),
			"solace_custom_attribute_definition": tableCustomAttributeDefinition(ctx),
			"solace_event_management_agent":      tableEventManagementAgent(ctx),
		},
	}

	return p
}

// NOT IMPLEMENTED
// "solace_datacenter_event_broker_service_version": tableDatacenterEventBrokerServiceVersion(ctx),
