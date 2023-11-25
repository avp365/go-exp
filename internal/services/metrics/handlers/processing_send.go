package handlers

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricName = "metric_name"
	handleName = "handle_name"
)

type ProcessingSend struct {
	mu     *sync.Mutex
	metric *prometheus.GaugeVec
	values map[string]CountRecord
}

type CountRecord struct {
	Total int
}

func (c *CountRecord) Add() {
	c.Total++
}
func NewProcessingSend() *ProcessingSend {
	return &ProcessingSend{
		mu: &sync.Mutex{},
		metric: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "MyMetrics",
			Help: "Handler send",
		}, []string{handleName, metricName}),
		values: make(map[string]CountRecord),
	}
}

func (m *ProcessingSend) Register(registry *prometheus.Registry) {
	registry.MustRegister(m.metric)
}

func (m *ProcessingSend) Add(message string) {

	m.mu.Lock()
	defer m.mu.Unlock()

	var record CountRecord

	if existRecord, ok := m.values[message]; ok {
		record = existRecord
	}

	record.Add()
	m.values[message] = record

}

func (m *ProcessingSend) SetToPrometheus() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.metric.Reset()

	for name, values := range m.values {

		m.metric.With(prometheus.Labels{handleName: name, metricName: "count"}).Set(float64(values.Total))
	}
	m.values = make(map[string]CountRecord)
}
