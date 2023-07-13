# Table: solace_configuration

Information about available configurations.

### Key columns
- Provide a numeric `id` if you want to query for a specific Configuration. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Configurations

```sql

select
  id, type
from
  solace_configuration;
```

### Detail of a Configuration

```sql
select
  *
from
  solace_configuration
where
  id = 'n5o4xx2fh62';
```
