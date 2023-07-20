# Table: solace_enum

In Event Portal, an enumeration is a bounded variable with a limited set of literal values.

### Key columns
- Provide a numeric `id` if you want to query for a specific enum. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all enums

select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  event_version_ref_count,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum;
```

### Details of an enum

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  event_version_ref_count,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum
where
  id = 'n5o4xx2fh62';
```

### List enums that have been created in the last 30 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  event_version_ref_count,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum
where
  created_time >= now() - interval '30' day;
```

### List enums that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  event_version_ref_count,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_enum
where
  updated_time <= now() - interval '10' day;
```