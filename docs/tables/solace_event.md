# Table: solace_event

In Event Portal, an event is an object that defines the properties that describe and categorize actual event instances.

### Key columns
- Provide a numeric `id` if you want to query for a specific event. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Events

```sql
select
  e.name as event,
  ev.version as version,
  ev.display_name as versionName
from
  solace_event e
  join
    solace_event_version ev
    on ev.event_id = e.id
where 
  e.id = '08ctmc2lyp6'

-- or a simplified version

select
  id, 
  name
from
  solace_event;
```

### Details of an Event

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
  solace_event
where
  id = 'n5o4xx2fh62';
```

### List Events that have been created in the last 30 days

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
  solace_event
where
  created_time >= now() - interval '30' day;
```

### List Events that have not been updated in the last 10 days

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
  solace_event
where
  updated_time <= now() - interval '10' day;
```