---
organization: Solace
category: ["saas"]
icon_url: "/images/plugins/SolaceLabs/solace.svg"
brand_color: "#00AD93"
display_name: "Solace PubSub+ Cloud"
short_name: "Solace"
description: "Solace PubSub+ Cloud plugin for exploring your Solace Cloud configuration in depth."
og_description: "Query Solace PubSub+ Cloud with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/SolaceLabs/solace-social-graphic.png"
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

- **[Table definitions & examples →](plugins/solacelabs/solace/tables)**

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

Installing the latest Solace PubSub+ Cloud plugin will create a config file (`~/.steampipe/config/solace.spc`) with a single connection named `solace`.

Uncomment and update the `api_token` value with the generated token (within quotes).

```hcl
connection "solace" {
  plugin = "SolaceLabs/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  # This can also be set via the `SOLACE_API_TOKEN` environment variable.
  # api_token = "XXXXXXXXX"

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  # This can also be set via the `SOLACE_API_URL` environment variable.
  api_url = "https://api.solace.cloud/"
}
```

Alternatively, you can also use the standard Solace environment variables to obtain credentials **only if other arguments (`api_token` and `api_url`) are not specified** in the connection:

```sh
export SOLACE_API_TOKEN=API_TOKEN
export SOLACE_API_URL=https://api.solace.cloud/
```

## Get involved

- Open source: https://github.com/solacelabs/steampipe-plugin-solace
- Solace PubSub+ Cloud REST API: https://api.solace.dev
- Ask the [Solace Community](https://solace.community)
- Community: [Slack Channel](https://steampipe.io/community/join)
