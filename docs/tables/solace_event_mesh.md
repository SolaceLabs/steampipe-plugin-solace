# Table: solace_event_mesh

Information about deployed Event Meshes on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Event Mesh. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Event Meshes

```sql

select
  id, 
  name,
  environment_id,
  description,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_mesh;
```

### Details of an Event Mesh

```sql
select
  id, 
  name,
  environment_id,
  description,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_mesh
where
  id = 'n5o4xx2fh62';
```

### List Event Meshes that have been created in the last 30 days

```sql
select
  id,
  name,
  application_type,
  broker_type,
  custom_attributes,
  created_time,
  created_by
from
  solace_event_mesh
where
  created_time >= now() - interval '30' day;
```

### List Event Meshes that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  broker_type,
  type,
  associated_entity_types,
  changed_by,
  updated_time
from
  solace_event_mesh
where
  updated_time <= now() - interval '10' day;
```