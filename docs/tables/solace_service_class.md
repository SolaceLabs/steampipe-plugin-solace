# Table: solace_service_class

Information about Service Classes on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Service Class. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Service Classes

```sql

select
  id, name
from
  solace_service_class;
```

### Detail of a Service Class

```sql
select
  *
from
  solace_service_class
where
  id = 'n5o4xx2fh62';
```