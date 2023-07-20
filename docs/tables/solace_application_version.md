# Table: solace_application_version

In Event Portal, when you update an Application, you can update an existing version or create a new version of the Application. Versions allow you to work on updates and test new versions in development and staging environments while the stable version remains in the production environment. Each version also has a lifecycle state. 

### Key columns
- Provide a numeric `id` if you want to query for a specific Application version. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all application versions

```sql
select
  av.version as version,
  av.display_name as versionName,
  a.name as application 
from
  solace_application_version av 
  join
    solace_application a 
    on av.application_id = a.id 
where
  av.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, 
  version, 
  display_name
from
  solace_application_version;
```

### Details of an application version

```sql
select
  id, 
  version, 
  display_name,
  declared_produced_event_version_ids,
  declared_consumed_event_version_ids,
  type,
  created_by,
  created_time
from
  solace_application_version
where
  id = 'n5o4xx2fh62';
```

### List application versions that have been created in the last 30 days

```sql
select
  id, 
  version, 
  display_name,
  declared_produced_event_version_ids,
  declared_consumed_event_version_ids,
  type,
  created_by,
  created_time
from
  solace_application_version
where
  created_time >= now() - interval '30' day;
```

### List application versions that have not been updated in the last 10 days

```sql
select
  id, 
  version, 
  display_name,
  declared_produced_event_version_ids,
  declared_consumed_event_version_ids,
  type,
  changed_by,
  updated_time
from
  solace_application_version
where
  updated_time <= now() - interval '10' day;
```