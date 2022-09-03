package metrics

import (
	"expvar"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

func NewMetricManager() *MetricManager {
	return &MetricManager{map[string]*Metric{}}
}

type MetricManager struct {
	metrics map[string]*Metric
}

func (mm *MetricManager) Register(metric *Metric) {
	// rewrite metrics
	mm.metrics[metric.Name] = metric
}

func (mm *MetricManager) RegisterMany(metrics []*Metric) {
	for _, metric := range metrics {
		mm.Register(metric)
	}
}

func (mm *MetricManager) Run(address string) {
	for name, c := range mm.metrics {
		expvar.Publish(name, c)
	}

	log.Info("metric server started")
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panic("metricManager::Run error while listen address", err)
	}
}
