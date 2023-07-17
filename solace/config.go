package solace

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {Type: schema.TypeString},
	"api_url":   {Type: schema.TypeString},
}

type Config struct {
	ApiToken *string `cty:"api_token"`
	ApiUrl   *string `cty:"api_url"`
}

func ConfigInstance() interface{} {
	return &Config{}
}
