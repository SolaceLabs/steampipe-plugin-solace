# Table: solace_messaging_service

Information about Messaging Services on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Messaging Service. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Messaging Services

```sql
select
  id,
  name
from
  solace_messaging_service;
```

### Details of a Messaging Service

```sql
select
  id,
  name,
  event_mesh_id,
  context_id,
  runtime_agent_id,
  created_time,
  created_by,
  changed_by,
  updated_time  
from
  solace_messaging_service
where
  id = 'n5o4xx2fh62';
```

### List Messaging Services that have been created in the last 30 days

```sql
select
  id,
  name,
  event_mesh_id,
  context_id,
  runtime_agent_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_messaging_service
where
  created_time >= now() - interval '30' day;
```

### List Messaging Services that have not been updated in the last 10 days

```sql
select
  id,
  name,
  event_mesh_id,
  context_id,
  runtime_agent_id,
  created_time,
  created_by,
  changed_by,
  updated_time  
from
  solace_messaging_service
where
  updated_time <= now() - interval '10' day;
```