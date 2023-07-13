# Table: solace_schema_version

In Event Portal, when you update a Schema, you can update an existing version or create a new version of the Schema. Versions allow you to work on updates and test new versions in development and staging environments while the stable version remains in the production environment. Each version also has a lifecycle state. 

### Key columns
- Provide a numeric `id` if you want to query for a specific Schema version. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Schema Versions

```sql
select
  sv.version as version,
  sv."displayName" as versionName,
  s.name as name
from
  solace_schema_version sv
  join
    solace_schema s
    on sv."enumId" = s.id
where sv.id = 'n5o41x2fh62';

-- or a simplified version

select
  id, version, displayName
from
  solace_schema_version;
```

### Detail of an Schema Version

```sql
select
  *
from
  solace_schema_version
where
  id = 'n5o4xx2fh62';
```
