package solace

import (
	"context"
	"encoding/json"
	"fmt"

	solace "github.com/SolaceLabs/steampipe-solace-go-client-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

var lastConnection *plugin.Connection
var solaceClient *solace.Client

func NewSolaceClient(c *plugin.Connection) (*solace.Client, error) {
	if lastConnection == c && solaceClient != nil {
		return solaceClient, nil
	}

	var cfg, ok = c.Config.(Config)
	if !ok {
		return nil, fmt.Errorf("config object is not valid")
	}

	var config, err = solace.NewConfig(cfg.ApiToken, cfg.ApiUrl)
	if err != nil {
		return nil, err
	}

	lastConnection = c
	solaceClient = solace.GetClient(config)

	return solaceClient, nil
}

func ToJSON(value interface{}) string {
	j, _ := json.Marshal(value)
	return string(j)
}

func LogQueryContext(namespace string, ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) {
	plugin.Logger(ctx).Info(namespace, "Table", d.Table.Name)
	plugin.Logger(ctx).Info(namespace, "QueryContext", ToJSON(d.QueryContext))
	plugin.Logger(ctx).Info(namespace, "EqualsQuals", ToJSON(d.EqualsQuals))
	plugin.Logger(ctx).Info(namespace, "HydrateData", ToJSON(h))
}

func StandardColumnDescription(key string) string {
	switch key {
	case "akas":
		return "Array of globally unique identifier strings (also known as) for the resource."
	case "tags":
		return "A map of tags for the resource."
	case "title":
		return "The display name for the resource."
	case "virtual":
		return "Virtual column, used to map the entity to another object."
	}
	return ""
}
