# Table: solace_enum

In Event Portal, an enumeration is a bounded variable with a limited set of literal values.

### Key columns
- Provide a numeric `id` if you want to query for a specific enum. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Enums

```sql
select
  e.name as event,
  ev.version as version,
  ev."displayName" as versionName
from
  solace_enum e
  join
    solace_enum_version ev
    on ev."enumId" = e.id
where 
  e.id = '08ctmc2lyp6'

-- or a simplified version

select
  id, name
from
  solace_enum;
```

### Detail of an Enum

```sql
select
  *
from
  solace_enum
where
  id = 'n5o4xx2fh62';
```
