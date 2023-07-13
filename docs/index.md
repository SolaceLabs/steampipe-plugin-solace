---
organization: Solace
category: ["saas"]
icon_url: "docs/images/solace.svg"
brand_color: "#6D01CC"
display_name: "Solace"
short_name: "solace"
description: "Solace plugin for exploring your Solace Cloud in depth."
og_description: "Query Solace Cloud with SQL! Open source CLI. No DB required."
og_image: "docs/images/solace-social-graphic.png"
---

# Solace Cloud + Steampipe

[Solace PubSub+](https://www.solace.com) is a cloud-based messaging and event streaming service provided by Solace. It offers a scalable and robust messaging infrastructure that enables real-time data movement and event-driven architecture in cloud and hybrid cloud environments.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.



For example:

```sql
select
  id,
  name,
  stats
from
  solace_application_domain;
```

```
+-------------+-------------------------------------+-------------------------------------------------------------------------------------------------------------------+
| id          | name                                | stats                                                                                                             |
+-------------+-------------------------------------+-------------------------------------------------------------------------------------------------------------------+
| n8xj0k6rx5i | AcmeRetail - Enterprise Governance  | {"applicationCount":0,"enumCount":3,"eventApiCount":1,"eventApiProductCount":0,"eventCount":2,"schemaCount":2}    |
| sfxq3pd8xcw | AcmeRetail - SAP S/4                | {"applicationCount":1,"enumCount":2,"eventApiCount":1,"eventApiProductCount":1,"eventCount":8,"schemaCount":26}   |
| dux1k1p9xsg | Acme Retail - CRM                   | {"applicationCount":1,"enumCount":0,"eventApiCount":0,"eventApiProductCount":0,"eventCount":3,"schemaCount":1}    |
| u2x73phaxbj | AcmeRetail - Human Relationships    | {"applicationCount":2,"enumCount":2,"eventApiCount":0,"eventApiProductCount":0,"eventCount":3,"schemaCount":5}    |
| 9nxoj6yfxm3 | AcmeRetail - Store Operations       | {"applicationCount":6,"enumCount":3,"eventApiCount":1,"eventApiProductCount":0,"eventCount":2,"schemaCount":2}    |
| 4tx0jilaxt2 | AcmeRetail - Supply Chain           | {"applicationCount":2,"enumCount":3,"eventApiCount":1,"eventApiProductCount":0,"eventCount":2,"schemaCount":2}    |
| 72x10oegx7u | AcmeRetailFacilities                | {"applicationCount":1,"enumCount":4,"eventApiCount":1,"eventApiProductCount":1,"eventCount":2,"schemaCount":2}    |
+-------------+-------------------------------------+-------------------------------------------------------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/solacelabs/solace/tables)**

## Get started

### Install

Download and install the latest Solace PubSub+ Cloud plugin:

```bash
steampipe plugin install solacelabs/solace
```

### Credentials

Solace PubSub+ Cloud plugin requires an API Token.

You can find more details on how to generate an API Token here - [Get your API token](https://docs.solace.com/Cloud/ght_api_tokens.htm). The generated token should be set as value for `api_token` parameter in the plugin config file (`~/.steampipe/config/solace.spc`) 

### Configuration

Installing the latest Solace PubSub+ Cloud plugin will create a config file (`~/.steampipe/config/solace.spc`) with a single connection named `solace`:

```hcl
connection "solace" {
  plugin = "local/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  # Steampipe will resolve the API key in below order:
  #   1. The "api_token" specified here in the config
  #   2. The `SOLACE_CLOUD_REST_API_TOKEN` environment variable
  api_token = ""

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  # Steampipe will resolve the API key in below order:
  #   1. The "api_url" specified here in the config
  #   2. The `SOLACE_CLOUD_REST_API_URL` environment variable
  api_url = "https://api.solace.cloud/"

  # Rate limiting
  # Solace Cloud REST API limits the number of requests you can send to the Cloud REST API API.
  # Solace Cloud REST API sets the rate limits based on your organization plan:
  # - Core: 60 per minute
  # - Pro: 120 per minute
  # - Teams: 240 per minute
  # - Enterprise: 1 000 per minute
  # We recommend to set a value below (or at most at) 80% of your total limit.
  # The default value is 50 if you don't override it here.
  # rate_limit = 50
}
```

## Get involved

- Open source: https://github.com/solacelabs/steampipe-plugin-solace
- Solace PubSub+ Cloud REST API: https://api.solace.dev
- Ask the [Solace Community](https://solace.community)
