# Table: solace_eventapi_version

In Event Portal, when you update an Event API, you can update an existing version or create a new version of the Event API. Versions allow you to work on updates and test new versions in development and staging environments while the stable version remains in the production environment. Each version also has a lifecycle state. 

### Key columns
- Provide a numeric `id` if you want to query for a specific Event API version. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Event API Versions

```sql
select
  ev.version as version,
  ev.display_name as versionName,
  e.name as name 
from
  solace_eventapi_version ev 
  join
    solace_eventapi e 
    on ev.event_api_id = e.id 
where
  ev.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, 
  version, 
  display_name
from
  solace_eventapi_version;
```

### Details of an Event API Version

```sql
select
  id, 
  version, 
  display_name,
  produced_event_version_ids,
  consumed_event_version_ids,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time   
from
  solace_eventapi_version
where
  id = 'n5o4xx2fh62';
```

### List Event API Versions that have been created in the last 30 days

```sql
select
  id, 
  version, 
  display_name,
  produced_event_version_ids,
  consumed_event_version_ids,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi_version
where
  created_time >= now() - interval '30' day;
```

### List Event API Versions that have not been updated in the last 10 days

```sql
select
  id, 
  version, 
  display_name,
  produced_event_version_ids,
  consumed_event_version_ids,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi_version
where
  updated_time <= now() - interval '10' day;
```