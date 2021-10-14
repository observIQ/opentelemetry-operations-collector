module github.com/GoogleCloudPlatform/opentelemetry-operations-collector

go 1.14

require (
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/mitchellh/go-testing-interface v1.0.3 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver v0.0.0-20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusexecreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver v0.36.1-0.20211013182413-e273899b2c27
	github.com/shirou/gopsutil v3.21.9+incompatible
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.36.1-0.20211004155959-190f8fbb2b9a
	go.opentelemetry.io/collector/model v0.36.1-0.20211004155959-190f8fbb2b9a
	go.uber.org/zap v1.19.1
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678
)
