# Table: solace_eventapi

In Event Portal, an event API bundles together the following information:

- Events you want to provide to other application developers
- Associated schema information
- Information about the event operation (publish and/or subscribe)


### Key columns
- Provide a numeric `id` if you want to query for a specific Event API. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Event APIs

```sql
select
  e.name as event,
  ev.version as version,
  ev.display_name as versionName
from
  solace_eventapi e
  join
    solace_eventapi_version ev
    on ev.event_api_id = e.id
where 
  e.id = '08ctmc2lyp6';

-- or a simplified version

select
  id,
  name,
  shared,
  application_domain_id,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time  
from
  solace_eventapi;
```

### Details of an Event API

```sql
select
  id,
  name,
  shared,
  application_domain_id,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time  
from
  solace_eventapi
where
  id = 'n5o4xx2fh62';
```

### List Event APIs that have been created in the last 30 days

```sql
select
  id,
  name,
  shared,
  application_domain_id,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time  
from
  solace_eventapi
where
  created_time >= now() - interval '30' day;
```

### List Event APIs that have not been updated in the last 10 days

```sql
select
  id,
  name,
  shared,
  application_domain_id,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi
where
  updated_time <= now() - interval '10' day;
```