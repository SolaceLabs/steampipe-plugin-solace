# Table: solace_custom_attribute_definition

Information about configured Custom Attributes on Event Portal resources.

### Key columns
- Provide a numeric `id` if you want to query for a specific Custom Attribute. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all custom attributes

```sql
select
  id, 
  name,
  value_type,
  scope,
  associated_entity_types,
  associated_entities
from
  solace_custom_attribute_definition;
```

### Details of a custom attribute

```sql
select
  id, 
  name,
  value_type,
  scope,
  associated_entity_types,
  associated_entities,
  created_by,
  created_time,
  changed_by,
  updated_time
from
  solace_custom_attribute_definition
where
  id = 'n5o4xx2fh62';
```

### List custom attributes that have been created in the last 30 days

```sql
select
  id, 
  name,
  value_type,
  scope,
  associated_entity_types,
  associated_entities,
  created_by,
  created_time,
  changed_by,
  updated_time
from
  solace_custom_attribute_definition
where
  created_time >= now() - interval '30' day;
```

### List custom attributes that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  value_type,
  scope,
  associated_entity_types,
  associated_entities,
  created_by,
  created_time,
  changed_by,
  updated_time
from
  solace_custom_attribute_definition
where
  updated_time <= now() - interval '10' day;
```