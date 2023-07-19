# Table: solace_configuration_type

Information about available configuration types.

### Key columns
- Provide a numeric `id` if you want to query for a specific configuration type. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all configuration type

```sql
select
  id, 
  name
from
  solace_configuration_type;
```

### Details of a configuration type

```sql
select
  id, 
  name,
  broker_type,
  type,
  associated_entity_types,
  created_time,
  created_by
from
  solace_configuration_type
where
  id = 'n5o4xx2fh62';
```

### List the configuration types that have been created in the last 30 days

```sql
select
  id, 
  name,
  broker_type,
  type,
  associated_entity_types,
  created_time,
  created_by
from
  solace_configuration_type
where
  created_time >= now() - interval '30' day;
```

### List the configuration types that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  broker_type,
  type,
  associated_entity_types,
  created_time,
  created_by
from
  solace_configuration_type
where
  updated_time <= now() - interval '10' day;
```