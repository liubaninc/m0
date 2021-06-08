package cmd

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "m0d"
)

// Metrics contains metrics exposed by this package.
// see MetricsProvider for descriptions.
type Metrics struct {
	MemUsage  metrics.Gauge
	CPUUsage  metrics.Gauge
	DiskUsage metrics.Gauge
	DiskTotal metrics.Gauge
}

// PrometheusMetrics returns Metrics build using Prometheus client library.
// Optionally, labels can be provided along with their values ("foo",
// "fooValue").
func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}

	return &Metrics{
		MemUsage: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "node_mem_usage",
			Help:      "mem usage rate of node.",
		}, labels).With(labelsAndValues...),
		DiskUsage: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "node_disk_usage",
			Help:      "disk usage rate of node.",
		}, labels).With(labelsAndValues...),
		DiskTotal: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "node_disk_total",
			Help:      "total disk of node.",
		}, labels).With(labelsAndValues...),
		CPUUsage: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "node_cpu_usage",
			Help:      "cpu usage rate of node.",
		}, labels).With(labelsAndValues...),
	}
}

// NopMetrics returns no-op Metrics.
func NopMetrics() *Metrics {
	return &Metrics{
		MemUsage:  discard.NewGauge(),
		DiskUsage: discard.NewGauge(),
		DiskTotal: discard.NewGauge(),
		CPUUsage:  discard.NewGauge(),
	}
}
