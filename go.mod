module github.com/GoogleCloudPlatform/opentelemetry-operations-collector

go 1.14

require (
	github.com/mitchellh/go-testing-interface v1.0.3 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpdreceiver v0.0.0-20210622143655-cc2154632229
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusexecreceiver v0.28.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver v0.28.0
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.28.0
	go.uber.org/zap v1.17.0
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpdreceiver => github.com/observIQ/opentelemetry-collector-contrib/receiver/httpdreceiver v0.0.0-20210622143655-cc2154632229
