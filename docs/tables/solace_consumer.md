# Table: solace_consumer

Information about configured consumers on an application version.

### Key columns
- Provide a numeric `id` if you want to query for a specific Consumer. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Consumers

```sql

select
  id, name
from
  solace_consumer;
```

### Detail of a Consumer

```sql
select
  *
from
  solace_consumer
where
  id = 'n5o4xx2fh62';
```
