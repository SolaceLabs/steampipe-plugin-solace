# Table: solace_datacenter

Information about available Data Centers on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Data Center. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List all data centers

```sql
select
  id, 
  name, 
  location,
  datacenter_type,
  provider,
  oper_state
from
  solace_datacenter;
```

### Details of a data center

```sql
select
  id, 
  name, 
  location,
  datacenter_type,
  provider,
  oper_state,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_datacenter
where
  id = 'n5o4xx2fh62';
```

### List data centers that have been created in the last 30 days

```sql
select
  id, 
  name, 
  location,
  datacenter_type,
  provider,
  oper_state,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_datacenter
where
  created_time >= now() - interval '30' day;
```

### List data centers that have not been updated in the last 10 days

```sql
select
  id, 
  name, 
  location,
  datacenter_type,
  provider,
  oper_state,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_datacenter
where
  updated_time <= now() - interval '10' day;
```