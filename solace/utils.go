package solace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

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

	apiToken := os.Getenv("SOLACE_API_TOKEN")
	apiURL := os.Getenv("SOLACE_API_URL")

	if cfg.ApiToken != nil {
		apiToken = *cfg.ApiToken
	}

	if cfg.ApiUrl != nil {
		apiURL = *cfg.ApiToken
	}

	if apiToken == "" {
		// Credentials not set
		return nil, errors.New("api_token must be configured")
	}

	if apiURL == "" {
		// Credentials not set
		return nil, errors.New("api_url must be configured")
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
	plugin.Logger(ctx).Debug(namespace, "Table", d.Table.Name)
	plugin.Logger(ctx).Debug(namespace, "QueryContext", ToJSON(d.QueryContext))
	plugin.Logger(ctx).Debug(namespace, "EqualsQuals", ToJSON(d.EqualsQuals))
	plugin.Logger(ctx).Debug(namespace, "HydrateData", ToJSON(h))
}
