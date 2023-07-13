# Table: solace_event_mesh

Information about deployed Event Meshes on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Event Mesh. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Event Meshes

```sql

select
  id, name
from
  solace_event_mesh;
```

### Detail of an Event Mesh

```sql
select
  *
from
  solace_event_mesh
where
  id = 'n5o4xx2fh62';
```
