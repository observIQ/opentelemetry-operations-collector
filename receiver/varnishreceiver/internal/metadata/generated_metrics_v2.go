// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for varnishreceiver metrics.
type MetricsSettings struct {
	VarnishBackendConnectionCount  MetricSettings `mapstructure:"varnish.backend.connection.count"`
	VarnishBackendRequestCount     MetricSettings `mapstructure:"varnish.backend.request.count"`
	VarnishCacheOperationCount     MetricSettings `mapstructure:"varnish.cache.operation.count"`
	VarnishClientRequestCount      MetricSettings `mapstructure:"varnish.client.request.count"`
	VarnishClientRequestErrorCount MetricSettings `mapstructure:"varnish.client.request.error.count"`
	VarnishObjectCount             MetricSettings `mapstructure:"varnish.object.count"`
	VarnishObjectExpired           MetricSettings `mapstructure:"varnish.object.expired"`
	VarnishObjectMoved             MetricSettings `mapstructure:"varnish.object.moved"`
	VarnishObjectNuked             MetricSettings `mapstructure:"varnish.object.nuked"`
	VarnishSessionCount            MetricSettings `mapstructure:"varnish.session.count"`
	VarnishThreadOperationCount    MetricSettings `mapstructure:"varnish.thread.operation.count"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		VarnishBackendConnectionCount: MetricSettings{
			Enabled: true,
		},
		VarnishBackendRequestCount: MetricSettings{
			Enabled: true,
		},
		VarnishCacheOperationCount: MetricSettings{
			Enabled: true,
		},
		VarnishClientRequestCount: MetricSettings{
			Enabled: true,
		},
		VarnishClientRequestErrorCount: MetricSettings{
			Enabled: true,
		},
		VarnishObjectCount: MetricSettings{
			Enabled: true,
		},
		VarnishObjectExpired: MetricSettings{
			Enabled: true,
		},
		VarnishObjectMoved: MetricSettings{
			Enabled: true,
		},
		VarnishObjectNuked: MetricSettings{
			Enabled: true,
		},
		VarnishSessionCount: MetricSettings{
			Enabled: true,
		},
		VarnishThreadOperationCount: MetricSettings{
			Enabled: true,
		},
	}
}

type metricVarnishBackendConnectionCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.backend.connection.count metric with initial data.
func (m *metricVarnishBackendConnectionCount) init() {
	m.data.SetName("varnish.backend.connection.count")
	m.data.SetDescription("The backend connection type count.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishBackendConnectionCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, backendConnectionTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.BackendConnectionType, pdata.NewAttributeValueString(backendConnectionTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishBackendConnectionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishBackendConnectionCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishBackendConnectionCount(settings MetricSettings) metricVarnishBackendConnectionCount {
	m := metricVarnishBackendConnectionCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishBackendRequestCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.backend.request.count metric with initial data.
func (m *metricVarnishBackendRequestCount) init() {
	m.data.SetName("varnish.backend.request.count")
	m.data.SetDescription("The backend requests count.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricVarnishBackendRequestCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishBackendRequestCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishBackendRequestCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishBackendRequestCount(settings MetricSettings) metricVarnishBackendRequestCount {
	m := metricVarnishBackendRequestCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishCacheOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.cache.operation.count metric with initial data.
func (m *metricVarnishCacheOperationCount) init() {
	m.data.SetName("varnish.cache.operation.count")
	m.data.SetDescription("The cache operation type count.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishCacheOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, cacheOperationsAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.CacheOperations, pdata.NewAttributeValueString(cacheOperationsAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishCacheOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishCacheOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishCacheOperationCount(settings MetricSettings) metricVarnishCacheOperationCount {
	m := metricVarnishCacheOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishClientRequestCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.client.request.count metric with initial data.
func (m *metricVarnishClientRequestCount) init() {
	m.data.SetName("varnish.client.request.count")
	m.data.SetDescription("The client request count.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishClientRequestCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.State, pdata.NewAttributeValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishClientRequestCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishClientRequestCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishClientRequestCount(settings MetricSettings) metricVarnishClientRequestCount {
	m := metricVarnishClientRequestCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishClientRequestErrorCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.client.request.error.count metric with initial data.
func (m *metricVarnishClientRequestErrorCount) init() {
	m.data.SetName("varnish.client.request.error.count")
	m.data.SetDescription("The client request errors received by status code.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishClientRequestErrorCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.HTTPStatusCode, pdata.NewAttributeValueString(httpStatusCodeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishClientRequestErrorCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishClientRequestErrorCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishClientRequestErrorCount(settings MetricSettings) metricVarnishClientRequestErrorCount {
	m := metricVarnishClientRequestErrorCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishObjectCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.object.count metric with initial data.
func (m *metricVarnishObjectCount) init() {
	m.data.SetName("varnish.object.count")
	m.data.SetDescription("The HTTP objects in the cache count.")
	m.data.SetUnit("{objects}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricVarnishObjectCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishObjectCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishObjectCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishObjectCount(settings MetricSettings) metricVarnishObjectCount {
	m := metricVarnishObjectCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishObjectExpired struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.object.expired metric with initial data.
func (m *metricVarnishObjectExpired) init() {
	m.data.SetName("varnish.object.expired")
	m.data.SetDescription("The expired objects from old age count.")
	m.data.SetUnit("{objects}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricVarnishObjectExpired) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishObjectExpired) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishObjectExpired) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishObjectExpired(settings MetricSettings) metricVarnishObjectExpired {
	m := metricVarnishObjectExpired{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishObjectMoved struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.object.moved metric with initial data.
func (m *metricVarnishObjectMoved) init() {
	m.data.SetName("varnish.object.moved")
	m.data.SetDescription("The moved operations done on the LRU list count.")
	m.data.SetUnit("{objects}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricVarnishObjectMoved) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishObjectMoved) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishObjectMoved) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishObjectMoved(settings MetricSettings) metricVarnishObjectMoved {
	m := metricVarnishObjectMoved{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishObjectNuked struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.object.nuked metric with initial data.
func (m *metricVarnishObjectNuked) init() {
	m.data.SetName("varnish.object.nuked")
	m.data.SetDescription("The objects that have been forcefully evicted from storage count.")
	m.data.SetUnit("{objects}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricVarnishObjectNuked) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishObjectNuked) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishObjectNuked) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishObjectNuked(settings MetricSettings) metricVarnishObjectNuked {
	m := metricVarnishObjectNuked{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishSessionCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.session.count metric with initial data.
func (m *metricVarnishSessionCount) init() {
	m.data.SetName("varnish.session.count")
	m.data.SetDescription("The session connection type count.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishSessionCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, sessionTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.SessionType, pdata.NewAttributeValueString(sessionTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishSessionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishSessionCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishSessionCount(settings MetricSettings) metricVarnishSessionCount {
	m := metricVarnishSessionCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricVarnishThreadOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills varnish.thread.operation.count metric with initial data.
func (m *metricVarnishThreadOperationCount) init() {
	m.data.SetName("varnish.thread.operation.count")
	m.data.SetDescription("The thread operation type count.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVarnishThreadOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, threadOperationsAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.ThreadOperations, pdata.NewAttributeValueString(threadOperationsAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVarnishThreadOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVarnishThreadOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVarnishThreadOperationCount(settings MetricSettings) metricVarnishThreadOperationCount {
	m := metricVarnishThreadOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                            pdata.Timestamp
	metricVarnishBackendConnectionCount  metricVarnishBackendConnectionCount
	metricVarnishBackendRequestCount     metricVarnishBackendRequestCount
	metricVarnishCacheOperationCount     metricVarnishCacheOperationCount
	metricVarnishClientRequestCount      metricVarnishClientRequestCount
	metricVarnishClientRequestErrorCount metricVarnishClientRequestErrorCount
	metricVarnishObjectCount             metricVarnishObjectCount
	metricVarnishObjectExpired           metricVarnishObjectExpired
	metricVarnishObjectMoved             metricVarnishObjectMoved
	metricVarnishObjectNuked             metricVarnishObjectNuked
	metricVarnishSessionCount            metricVarnishSessionCount
	metricVarnishThreadOperationCount    metricVarnishThreadOperationCount
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                            pdata.NewTimestampFromTime(time.Now()),
		metricVarnishBackendConnectionCount:  newMetricVarnishBackendConnectionCount(settings.VarnishBackendConnectionCount),
		metricVarnishBackendRequestCount:     newMetricVarnishBackendRequestCount(settings.VarnishBackendRequestCount),
		metricVarnishCacheOperationCount:     newMetricVarnishCacheOperationCount(settings.VarnishCacheOperationCount),
		metricVarnishClientRequestCount:      newMetricVarnishClientRequestCount(settings.VarnishClientRequestCount),
		metricVarnishClientRequestErrorCount: newMetricVarnishClientRequestErrorCount(settings.VarnishClientRequestErrorCount),
		metricVarnishObjectCount:             newMetricVarnishObjectCount(settings.VarnishObjectCount),
		metricVarnishObjectExpired:           newMetricVarnishObjectExpired(settings.VarnishObjectExpired),
		metricVarnishObjectMoved:             newMetricVarnishObjectMoved(settings.VarnishObjectMoved),
		metricVarnishObjectNuked:             newMetricVarnishObjectNuked(settings.VarnishObjectNuked),
		metricVarnishSessionCount:            newMetricVarnishSessionCount(settings.VarnishSessionCount),
		metricVarnishThreadOperationCount:    newMetricVarnishThreadOperationCount(settings.VarnishThreadOperationCount),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricVarnishBackendConnectionCount.emit(metrics)
	mb.metricVarnishBackendRequestCount.emit(metrics)
	mb.metricVarnishCacheOperationCount.emit(metrics)
	mb.metricVarnishClientRequestCount.emit(metrics)
	mb.metricVarnishClientRequestErrorCount.emit(metrics)
	mb.metricVarnishObjectCount.emit(metrics)
	mb.metricVarnishObjectExpired.emit(metrics)
	mb.metricVarnishObjectMoved.emit(metrics)
	mb.metricVarnishObjectNuked.emit(metrics)
	mb.metricVarnishSessionCount.emit(metrics)
	mb.metricVarnishThreadOperationCount.emit(metrics)
}

// RecordVarnishBackendConnectionCountDataPoint adds a data point to varnish.backend.connection.count metric.
func (mb *MetricsBuilder) RecordVarnishBackendConnectionCountDataPoint(ts pdata.Timestamp, val int64, backendConnectionTypeAttributeValue string) {
	mb.metricVarnishBackendConnectionCount.recordDataPoint(mb.startTime, ts, val, backendConnectionTypeAttributeValue)
}

// RecordVarnishBackendRequestCountDataPoint adds a data point to varnish.backend.request.count metric.
func (mb *MetricsBuilder) RecordVarnishBackendRequestCountDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricVarnishBackendRequestCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordVarnishCacheOperationCountDataPoint adds a data point to varnish.cache.operation.count metric.
func (mb *MetricsBuilder) RecordVarnishCacheOperationCountDataPoint(ts pdata.Timestamp, val int64, cacheOperationsAttributeValue string) {
	mb.metricVarnishCacheOperationCount.recordDataPoint(mb.startTime, ts, val, cacheOperationsAttributeValue)
}

// RecordVarnishClientRequestCountDataPoint adds a data point to varnish.client.request.count metric.
func (mb *MetricsBuilder) RecordVarnishClientRequestCountDataPoint(ts pdata.Timestamp, val int64, stateAttributeValue string) {
	mb.metricVarnishClientRequestCount.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// RecordVarnishClientRequestErrorCountDataPoint adds a data point to varnish.client.request.error.count metric.
func (mb *MetricsBuilder) RecordVarnishClientRequestErrorCountDataPoint(ts pdata.Timestamp, val int64, httpStatusCodeAttributeValue string) {
	mb.metricVarnishClientRequestErrorCount.recordDataPoint(mb.startTime, ts, val, httpStatusCodeAttributeValue)
}

// RecordVarnishObjectCountDataPoint adds a data point to varnish.object.count metric.
func (mb *MetricsBuilder) RecordVarnishObjectCountDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricVarnishObjectCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordVarnishObjectExpiredDataPoint adds a data point to varnish.object.expired metric.
func (mb *MetricsBuilder) RecordVarnishObjectExpiredDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricVarnishObjectExpired.recordDataPoint(mb.startTime, ts, val)
}

// RecordVarnishObjectMovedDataPoint adds a data point to varnish.object.moved metric.
func (mb *MetricsBuilder) RecordVarnishObjectMovedDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricVarnishObjectMoved.recordDataPoint(mb.startTime, ts, val)
}

// RecordVarnishObjectNukedDataPoint adds a data point to varnish.object.nuked metric.
func (mb *MetricsBuilder) RecordVarnishObjectNukedDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricVarnishObjectNuked.recordDataPoint(mb.startTime, ts, val)
}

// RecordVarnishSessionCountDataPoint adds a data point to varnish.session.count metric.
func (mb *MetricsBuilder) RecordVarnishSessionCountDataPoint(ts pdata.Timestamp, val int64, sessionTypeAttributeValue string) {
	mb.metricVarnishSessionCount.recordDataPoint(mb.startTime, ts, val, sessionTypeAttributeValue)
}

// RecordVarnishThreadOperationCountDataPoint adds a data point to varnish.thread.operation.count metric.
func (mb *MetricsBuilder) RecordVarnishThreadOperationCountDataPoint(ts pdata.Timestamp, val int64, threadOperationsAttributeValue string) {
	mb.metricVarnishThreadOperationCount.recordDataPoint(mb.startTime, ts, val, threadOperationsAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// NewMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) NewMetricData() pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/varnishreceiver")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// BackendConnectionType (The backend connection types.)
	BackendConnectionType string
	// CacheName (The varnish cache name.)
	CacheName string
	// CacheOperations (The cache operation types)
	CacheOperations string
	// HTTPStatusCode (An HTTP status code.)
	HTTPStatusCode string
	// SessionType (The session connection types.)
	SessionType string
	// State (The client request states.)
	State string
	// ThreadOperations (The thread operation types.)
	ThreadOperations string
}{
	"kind",
	"cache_name",
	"operation",
	"status_code",
	"kind",
	"state",
	"operation",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeBackendConnectionType are the possible values that the attribute "backend_connection_type" can have.
var AttributeBackendConnectionType = struct {
	Success   string
	Recycle   string
	Reuse     string
	Fail      string
	Unhealthy string
	Busy      string
	Retry     string
}{
	"success",
	"recycle",
	"reuse",
	"fail",
	"unhealthy",
	"busy",
	"retry",
}

// AttributeCacheOperations are the possible values that the attribute "cache_operations" can have.
var AttributeCacheOperations = struct {
	Hit     string
	Miss    string
	HitPass string
}{
	"hit",
	"miss",
	"hit_pass",
}

// AttributeSessionType are the possible values that the attribute "session_type" can have.
var AttributeSessionType = struct {
	Accepted string
	Dropped  string
	Failed   string
}{
	"accepted",
	"dropped",
	"failed",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Received string
	Dropped  string
}{
	"received",
	"dropped",
}

// AttributeThreadOperations are the possible values that the attribute "thread_operations" can have.
var AttributeThreadOperations = struct {
	Created   string
	Destroyed string
	Failed    string
}{
	"created",
	"destroyed",
	"failed",
}