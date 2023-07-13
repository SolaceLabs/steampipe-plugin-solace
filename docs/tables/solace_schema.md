# Table: solace_enum

In Event Portal, a schema defines the payload of an event. Producers and consumers of an event can both understand an event's payload based on the schema definition assigned to the event. Some events need complex schemas to describe the properties of each value that could be used in the event.

### Key columns
- Provide a numeric `id` if you want to query for a specific schema. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Schemas

```sql
select
  s.name as schema,
  sv.version as version,
  sv."displayName" as versionName
from
  solace_schema s
  join
    solace_schema_version sv
    on sv."schemaId" = s.id
where 
  s.id = '08ctmc2lyp6'

-- or a simplified version

select
  id, name
from
  solace_schema;
```

### Detail of a Schema

```sql
select
  *
from
  solace_schema
where
  id = 'n5o4xx2fh62';
```
