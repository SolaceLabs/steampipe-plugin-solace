# Table: solace_environment

Information about available Environments on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Environment. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Environments

```sql

select
  id, name, description
from
  solace_environment;
```

### Detail of an Environment

```sql
select
  *
from
  solace_environment
where
  id = 'n5o4xx2fh62';
```
