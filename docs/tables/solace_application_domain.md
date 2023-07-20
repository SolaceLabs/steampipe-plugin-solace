# Table: solace_application_domain

An application domain represents a namespace where applications, events, and other EDA objects reside. Application domains organize your applications, events, and other associated objects for different teams, groups, or lines of business within your organization.

### Key columns
- Provide a numeric `id` if you want to query for a specific domain. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all application domains

```sql
select
  d.name as domain,
  a.name as application 
from
  solace_application_domain d 
  join
    solace_application a 
    on a.application_domain_id = d.id;

-- or a simplified version

select
  id, 
  name
from
  solace_application_domain;
```

### Details of an application domain

```sql
select
  id,
  name,
  application_type,
  broker_type,
  custom_attributes,
  created_time,
  created_by
from
  solace_application_domain
where
  id = 'e32trlalpa4';
```

### List application domains that have been created in the last 30 days

```sql
select
  id,
  name,
  description,
  stats,
  created_by,
  created_time
from
  solace_application_domain
where
  created_time >= now() - interval '30' day;
```

### List application domains that have not been updated in the last 30 days

```sql
select
  id,
  name,
  description,
  stats,
  changed_by,
  updated_time
from
  solace_application_domain
where
  updated_time <= now() - interval '10' day;
```
