# Table: solace_messaging_service

Information about Messaging Services on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Messaging Service. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Messaging Services

```sql

select
  id, name
from
  solace_messaging_service;
```

### Detail of a Messaging Service

```sql
select
  *
from
  solace_messaging_service
where
  id = 'n5o4xx2fh62';
```
