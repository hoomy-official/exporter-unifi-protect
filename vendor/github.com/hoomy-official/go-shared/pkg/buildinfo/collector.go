package buildinfo

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	buildInfoDesc *prometheus.Desc
}

func NewCollector(u BuildInfo) *Collector {
	return &Collector{
		buildInfoDesc: prometheus.NewDesc(
			"build_info",
			"Information about the binary build.",
			nil,
			prometheus.Labels{
				"name":        u.Name(),
				"version":     u.Version(),
				"date":        u.Date(),
				"buildSource": u.BuildSource(),
				"commit":      u.Commit(),
			},
		),
	}
}

// Describe returns all descriptions of the collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.buildInfoDesc
}

// Collect returns the current state of all metrics of the collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.buildInfoDesc, prometheus.GaugeValue, 1)
}
