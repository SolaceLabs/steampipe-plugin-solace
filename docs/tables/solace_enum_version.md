# Table: solace_enum_version

In Event Portal, when you update an Enum, you can update an existing version or create a new version of the Enum. Versions allow you to work on updates and test new versions in development and staging environments while the stable version remains in the production environment. Each version also has a lifecycle state. 

### Key columns
- Provide a numeric `id` if you want to query for a specific Enum version. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all enum versions

```sql
select
  ev.version as version,
  ev."displayName" as versionName,
  e.name as name
from
  solace_enum_version ev
  join
    solace_enum e
    on ev."enumId" = e.id
where ev.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, 
  version, 
  displayName
from
  solace_enum_version;
```

### Detail of an enum version

```sql
select
  id, 
  version, 
  displayName,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum_version
where
  id = 'n5o4xx2fh62';
```

### List applications that have been created in the last 30 days

```sql
select
  id, 
  version, 
  displayName,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum_version
where
  created_time >= now() - interval '30' day;
```

### List applications that have not been updated in the last 10 days

```sql
select
  id, 
  version, 
  displayName,
  state_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum_version
where
  updated_time <= now() - interval '10' day;
```