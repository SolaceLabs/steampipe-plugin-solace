# Table: solace_service_class

Information about Service Classes on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Service Class. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Service Classes

```sql
select
  id, 
  name,
  type
from
  solace_service_class;
```

### Details of a Service Class

```sql
select
  id, 
  name,
  type,
  vpn_connections,
  broker_scaling_tier,
  vpn_max_spool_size
from
  solace_service_class
where
  id = 'n5o4xx2fh62';
```

### List service classes which have more than 5 VPN connections

```sql
select
  id, 
  name,
  type,
  vpn_connections,
  broker_scaling_tier,
  vpn_max_spool_size
from
  solace_service_class
where
  vpn_connections > 5;
```