# Table: solace_custom_attribute_definition

Information about configured Custom Attributes on Event Portal resources.

### Key columns
- Provide a numeric `id` if you want to query for a specific Custom Attribute. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Custom Attributes

```sql

select
  id, name
from
  solace_custom_attribute_definition;
```

### Detail of a Custom Attribute

```sql
select
  *
from
  solace_custom_attribute_definition
where
  id = 'n5o4xx2fh62';
```
