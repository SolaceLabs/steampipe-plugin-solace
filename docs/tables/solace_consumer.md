# Table: solace_consumer

Information about configured consumers on an application version.

### Key columns
- Provide a numeric `id` if you want to query for a specific Consumer. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all consumers

```sql
select
  id, 
  name,
  consumer_type,
  broker_type,
  application_version_id,
  subscriptions,
  type
from
  solace_consumer;
```

### Details of a consumer

```sql
select
  id, 
  name,
  consumer_type,
  broker_type,
  application_version_id,
  subscriptions,
  type
from
  solace_consumer
where
  id = 'n5o4xx2fh62';
```

### List consumers that have been created in the last 30 days

```sql
select
  id, 
  name,
  consumer_type,
  broker_type,
  application_version_id,
  subscriptions,
  type
  created_time,
  created_by
from
  solace_consumer
where
  created_time >= now() - interval '30' day;
```

### List consumers that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  consumer_type,
  broker_type,
  application_version_id,
  subscriptions,
  type,
  changed_by,
  updated_time
from
  solace_consumer
where
  updated_time <= now() - interval '10' day;
```