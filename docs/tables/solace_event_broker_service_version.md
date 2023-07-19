# Table: solace_event_broker_service_version

Details of Broker Services Version on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Broker Service version. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List details of all Broker Service versions

```sql
select
  id, 
  version,
  supported_service_classes,
  capabilities
from
  solace_event_broker_service_version;
```

### Details of a Broker Service version

```sql
select
  id, 
  version,
  supported_service_classes,
  capabilities
from
  solace_event_broker_service_version
where
  id = 'n5o4xx2fh62';
```