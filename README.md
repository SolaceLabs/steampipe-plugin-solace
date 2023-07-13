![image](docs/images/steampipe-solace-plugin.png)

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

# Solace PubSub+ Cloud + Steampipe

## Overview
The goal is to retrieve information on [Solace PubSub+ Cloud](https://solace.com/products/platform/cloud/) resources via SQL queries using Solace plugin for Steampipe.

[Steampipe](https://steampipe.io) is an open-source framework for querying, analyzing, and working with data from various sources. It provides a unified and declarative way to interact with disparate data sources, such as databases, cloud providers, and APIs, using SQL syntax.

Plugins are what Steampipe uses to define the schema for remote resources. The Solace plugin for Steampipe defines which tables are available and performs API calls to query those resources.

3. Run a query

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
| dux1k1p9xsg | AcmeRetail - CRM                    | {"applicationCount":1,"enumCount":0,"eventApiCount":0,"eventApiProductCount":0,"eventCount":3,"schemaCount":1}    |
| u2x73phaxbj | AcmeRetail - Human Relationships    | {"applicationCount":2,"enumCount":2,"eventApiCount":0,"eventApiProductCount":0,"eventCount":3,"schemaCount":5}    |
| 9nxoj6yfxm3 | AcmeRetail - Store Operations       | {"applicationCount":6,"enumCount":3,"eventApiCount":1,"eventApiProductCount":0,"eventCount":2,"schemaCount":2}    |
| 4tx0jilaxt2 | AcmeRetail - Supply Chain           | {"applicationCount":2,"enumCount":3,"eventApiCount":1,"eventApiProductCount":0,"eventCount":2,"schemaCount":2}    |
| 72x10oegx7u | AcmeRetail - Facilities             | {"applicationCount":1,"enumCount":4,"eventApiCount":1,"eventApiProductCount":1,"eventCount":2,"schemaCount":2}    |
+-------------+-------------------------------------+-------------------------------------------------------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](docs/index.md)**


## Getting started quickly
1. Install the plugin with [Steampipe](https://steampipe.io/):
```shell
steampipe plugin install solacelabs/solace
```
2. Configure your credentials and config file.

Configure your [credentials](https://hub.steampipe.io/plugins/solacelabs/solace#credentials) and [config file](https://hub.steampipe.io/plugins/solacelabs/solace/make#configuration).

## Resources
This is not an officially supported Solace product.

For more information try these resources:
- Ask the [Solace Community](https://solace.community)
- The Solace Developer Portal website at: https://solace.dev


## Contributing
Contributions are encouraged! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors
See the list of [contributors](https://github.com/solacecommunity/<github-repo>/graphs/contributors) who participated in this project.

## License
See the [LICENSE](LICENSE) file for details.
