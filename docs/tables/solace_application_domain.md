# Table: solace_application_domain

An application domain represents a namespace where applications, events, and other EDA objects reside. Application domains organize your applications, events, and other associated objects for different teams, groups, or lines of business within your organization.

### Key columns
- Provide a numeric `id` if you want to query for a specific domain. This can be either set directly in a `where` clause, or specified as part of `join` with another table.

### Caveat
- Be careful when requesting all columns (`*`) without using an `id` in the query. To load this data, Steampipe will have to issue multiple API request to retrieve all resources (essentially issuing a paginated queries).

## Examples

### List of all Application Domains

```sql
select
  d.name as domain,
  a.name as application
from
  solace_application_domain d
  join
    solace_application a
    on a."applicationDomainId" = d.id

-- or a simplified version

select
  id, name
from
  solace_application_domain;
```

### Detail of an Application Domain

```sql
select
  *
from
  solace_application_domain
where
  id = 'e32trlalpa4';
```
