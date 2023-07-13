# Table: solace_configuration_type

Information about available configuration types.

### Key columns
- Provide a numeric `id` if you want to query for a specific configuration type. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Configuration Type

```sql

select
  id, name
from
  solace_configuration_type;
```

### Detail of an Configuration Type

```sql
select
  *
from
  solace_configuration_type
where
  id = 'n5o4xx2fh62';
```
