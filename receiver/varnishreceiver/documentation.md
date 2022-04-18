[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# varnishreceiver

## Metrics

These are the metrics available for this scraper.

| Name | Description | Unit | Type | Attributes |
| ---- | ----------- | ---- | ---- | ---------- |
| **varnish.backend.connection.count** | The backend connection type count. | {connections} | Sum(Int) | <ul> <li>backend_connection_type</li> </ul> |
| **varnish.backend.request.count** | The backend requests count. | {requests} | Sum(Int) | <ul> </ul> |
| **varnish.cache.operation.count** | The cache operation type count. | {operations} | Sum(Int) | <ul> <li>cache_operations</li> </ul> |
| **varnish.client.request.count** | The client request count. | {requests} | Sum(Int) | <ul> <li>state</li> </ul> |
| **varnish.client.request.error.count** | The client request errors received by status code. | {requests} | Sum(Int) | <ul> <li>http.status_code</li> </ul> |
| **varnish.object.count** | The HTTP objects in the cache count. | {objects} | Sum(Int) | <ul> </ul> |
| **varnish.object.expired** | The expired objects from old age count. | {objects} | Sum(Int) | <ul> </ul> |
| **varnish.object.moved** | The moved operations done on the LRU list count. | {objects} | Sum(Int) | <ul> </ul> |
| **varnish.object.nuked** | The objects that have been forcefully evicted from storage count. | {objects} | Sum(Int) | <ul> </ul> |
| **varnish.session.count** | The session connection type count. | {connections} | Sum(Int) | <ul> <li>session_type</li> </ul> |
| **varnish.thread.operation.count** | The thread operation type count. | {operations} | Sum(Int) | <ul> <li>thread_operations</li> </ul> |

**Highlighted metrics** are emitted by default. Other metrics are optional and not emitted by default.
Any metric can be enabled or disabled with the following scraper configuration:

```yaml
metrics:
  <metric_name>:
    enabled: <true|false>
```

## Attributes

| Name | Description |
| ---- | ----------- |
| backend_connection_type | The backend connection types. |
| cache_name | The varnish cache name. |
| cache_operations | The cache operation types |
| http.status_code | An HTTP status code. |
| session_type | The session connection types. |
| state | The client request states. |
| thread_operations | The thread operation types. |