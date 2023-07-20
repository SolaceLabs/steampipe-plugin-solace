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

- **[Table definitions & examples →](tables/index.md)**

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

Update the `api_token` value with the generated token (within quotes).

```hcl
connection "solace" {
  plugin = "SolaceLabs/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  # This can also be set via the `SOLACE_API_TOKEN` environment variable.
  # api_token = "eABCGciOiJSUzI1NiIsImtpMQITm1hYXNfcHJvZF8yMDIwMDMyNiIsInR5cCI6Ikp123J9.eyJvcmciOiJ5aG11aWU1azB5cCIsIm9yZ1R5cGUiOiJFTlRFUlBSSVNFIiwic3ViIjoiMmtmMGZ6ZjVjaDciLCJwZXJtaXNzaW9ucyI6IkFBQUFBSUFQQUFBQWZ6Z0E0QUVBQUFBQUFBQUFBQUFBQUlDeHpvY2hJQWpnTC8vL2c1WGZCZDREV01NRDQ0ZS9NUT09IiwiYXBpVG9rZW5JZCI6ImNkcWdpeHo5bHNzIiwiaXNzIjoiU29sYWNlIENvcnBvcmF0aW9uIiwiaWF0IjoxNjg5Nzc0NjMxfQ.Ry_uuVdcb4A_eFmK8LNnpFq1234W3wBwDxdebal8iz5CVWC1M1P48ao4w64MH1234RXpm9ZewguzINXO8hrAbcAa41EqRWrRK5V1P2Rpa78_z2YyCaO98ah5TtPbWdpytga6BxKJsy2I2sawdD8PIm7wbJpty9e4UG6fZEVSXDpcVA2QWERTHvIPV5PwOuT7WDa-IvIDIDeQpD8Dy44QvavpcqDaQhafW5B_P4EU715RVYuepVDgFtMLVXOuyz2m5DVab6BuizQ6zoi9adb1hO46j0bRi9M-eyG3B20ho5N_h0Jhd7GrDdqIaxoLRoKbqu4QWn6GnjH4prGd5H1mrrPQR"

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  # This can also be set via the `SOLACE_API_URL` environment variable.
  api_url = "https://api.solace.cloud/"
}
```

Alternatively, you can also use the standard Solace environment variables to obtain credentials **only if other arguments (`apo_token` and `api_url`) are not specified** in the connection:

```sh
export SOLACE_API_TOKEN=eABCGciOiJSUzI1NiIsImtpMQITm1hYXNfcHJvZF8yMDIwMDMyNiIsInR5cCI6Ikp123J9
export SOLACE_API_URL=https://api.solace.cloud/
```

## Get involved

- Open source: https://github.com/solacelabs/steampipe-plugin-solace
- Solace PubSub+ Cloud REST API: https://api.solace.dev
- Ask the [Solace Community](https://solace.community)
- Community: [Slack Channel](https://steampipe.io/community/join)
