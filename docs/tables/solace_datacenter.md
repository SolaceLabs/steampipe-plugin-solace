# Table: solace_datacenter

Information about available Data Centers on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Data Center. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Data Centers

```sql

select
  id, name, location
from
  solace_datacenter;
```

### Detail of a Data Center

```sql
select
  *
from
  solace_datacenter
where
  id = 'n5o4xx2fh62';
```
