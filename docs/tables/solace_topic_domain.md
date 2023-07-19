# Table: solace_topic_domain

Information about Topic Domains on the Solace PubSub+ Cloud.

### Key columns
- Provide a numeric `id` if you want to query for a specific Topic Domain. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Topic Domains

```sql
select
  id, 
  type,
  application_domain_id,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_topic_domain;
```

### Details of a Topic Domain

```sql
select
  id, 
  type,
  application_domain_id,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_topic_domain
where
  id = 'n5o4xx2fh62';
```

### List Topic Domains that have been created in the last 30 days

```sql
select
  id, 
  type,
  application_domain_id,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_topic_domain
where
  created_time >= now() - interval '30' day;
```

### List Topic Domains that have not been updated in the last 10 days

```sql
select
  id, 
  type,
  application_domain_id,
  broker_type,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_topic_domain
where
  updated_time <= now() - interval '10' day;
```