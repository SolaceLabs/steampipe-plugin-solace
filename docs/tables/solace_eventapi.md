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
  ev."displayName" as versionName
from
  solace_eventapi e
  join
    solace_eventapi_version ev
    on ev."eventApiId" = e.id
where 
  e.id = '08ctmc2lyp6'

-- or a simplified version

select
  id, name
from
  solace_eventapi;
```

### Detail of an Event API

```sql
select
  *
from
  solace_eventapi
where
  id = 'n5o4xx2fh62';
```
