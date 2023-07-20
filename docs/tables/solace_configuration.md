# Table: solace_configuration

Information about available configurations.

### Key columns
- Provide a numeric `id` if you want to query for a specific Configuration. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all configurations

```sql
select
  id, 
  type,
  context_type,
  context_id,
  created_by,
  created_time
from
  solace_configuration;
```

### Details of a configuration

```sql
select
  id, 
  type,
  context_type,
  context_id,
  created_by,
  created_time
from
  solace_configuration
where
  id = 'n5o4xx2fh62';
```

### List configurations that have been created in the last 30 days

```sql
select
  id, 
  type,
  context_type,
  context_id,
  created_by,
  created_time
from
  solace_configuration
where
  created_time >= now() - interval '30' day;
```

### List configurations that have not been updated in the last 10 days

```sql
select
  id, 
  type,
  context_type,
  context_id,
  created_by,
  created_time,
  changed_by,
  updated_time
from
  solace_configuration
where
  updated_time <= now() - interval '10' day;
```