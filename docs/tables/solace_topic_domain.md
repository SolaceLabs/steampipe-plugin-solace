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
  id, type
from
  solace_topic_domain;
```

### Detail of a Topic Domain

```sql
select
  *
from
  solace_topic_domain
where
  id = 'n5o4xx2fh62';
```
