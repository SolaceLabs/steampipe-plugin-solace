# Table: solace_event_broker_service_detail

Details of Broker Services on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Broker Service. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List details of all Broker Services

```sql
select
  id, 
  name,
  type,
  owned_by,
  infrastructure_id,
  datacenter_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_broker_service_detail;
```

### Details of a Broker Service

```sql
select
  id, 
  name,
  type,
  owned_by,
  infrastructure_id,
  datacenter_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_broker_service_detail
where
  id = 'n5o4xx2fh62';
```

### List applications that have been created in the last 30 days

```sql
select
  id, 
  name,
  type,
  owned_by,
  infrastructure_id,
  datacenter_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_broker_service_detail
where
  created_time >= now() - interval '30' day;
```

### List applications that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  type,
  owned_by,
  infrastructure_id,
  datacenter_id,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_event_broker_service_detail
where
  updated_time <= now() - interval '10' day;
```