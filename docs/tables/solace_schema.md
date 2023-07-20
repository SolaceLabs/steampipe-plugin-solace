# Table: solace_enum

In Event Portal, a schema defines the payload of an event. Producers and consumers of an event can both understand an event's payload based on the schema definition assigned to the event. Some events need complex schemas to describe the properties of each value that could be used in the event.

### Key columns
- Provide a numeric `id` if you want to query for a specific schema. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Schemas

```sql
select
  s.name as schema,
  sv.version as version,
  sv.display_name as versionName
from
  solace_schema s
  join
    solace_schema_version sv
    on sv.schema_id = s.id
where 
  s.id = '08ctmc2lyp6';

-- or a simplified version

select
  id, 
  name
from
  solace_schema;
```

### Details of a Schema

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  schema_type,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time    
from
  solace_schema
where
  id = 'n5o4xx2fh62';
```

### List Schemas that have been created in the last 30 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  schema_type,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_schema
where
  created_time >= now() - interval '30' day;
```

### List Schemas that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  schema_type,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_schema
where
  updated_time <= now() - interval '10' day;
```