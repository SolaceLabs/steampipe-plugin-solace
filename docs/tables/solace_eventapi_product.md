# Table: solace_eventapi_product

In Event Portal, an Even API Product is a bundle of event APIs that you can provide to other organizations, so they can consume the events that you have included in your event APIs. An Event API Product contains one or more event APIs plus service plan details for deploying the product to allow others to consume the events.

### Key columns
- Provide a numeric `id` if you want to query for a specific Event API Product. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Event API Products

```sql
select
  e.name as event,
  ev.version as version,
  ev."displayName" as versionName
from
  solace_eventapi_product e
  join
    solace_eventapi_product_version ev
    on ev."eventApiProductId" = e.id
where 
  e.id = '08ctmc2lyp6';

-- or a simplified version

select
  id, 
  name
from
  solace_eventapi_product;
```

### Detail of an Event API Product

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi_product
where
  id = 'n5o4xx2fh62';
```

### List Event API Products that have been created in the last 30 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi_product
where
  created_time >= now() - interval '30' day;
```

### List Event API Products that have not been updated in the last 10 days

```sql
select
  id, 
  name,
  application_domain_id,
  shared,
  number_of_versions,
  created_time,
  created_by,
  changed_by,
  updated_time
from
  solace_eventapi_product
where
  updated_time <= now() - interval '10' day;
```