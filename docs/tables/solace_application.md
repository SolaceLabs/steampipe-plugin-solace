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
  av.display_name as versionName 
from
  solace_application a 
  join
    solace_application_version av 
    on av.application_id = a.id 
where
  a.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, 
  name
from
  solace_application;
```

### Details of an application

```sql
select
  id,
  name,
  application_type,
  broker_type,
  custom_attributes
from
  solace_application
where
  id = 'n5o4xx2fh62';
```

### List applications that have been created in the last 30 days

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
  solace_application
where
  created_time >= now() - interval '30' day;
```

### List applications that have not been updated in the last 10 days

```sql
select
  id,
  name,
  application_type,
  broker_type,
  custom_attributes,
  updated_time,
  changed_by
from
  solace_application
where
  updated_time <= now() - interval '10' day;
```