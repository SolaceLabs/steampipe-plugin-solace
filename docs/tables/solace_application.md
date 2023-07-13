# Table: solace_application

An application in Event Portal is an object that represents software that produces and consumes events

### Key columns
- Provide a numeric `id` if you want to query for a specific application. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Applications

```sql
select
  a.name as application,
  av.version as version,
  av."displayName" as versionName
from
  solace_application a
  join
    solace_application_version av
    on av."applicationId" = a.id
where a.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, name
from
  solace_application;
```

### Detail of an Application

```sql
select
  *
from
  solace_application
where
  id = 'n5o4xx2fh62';
```
