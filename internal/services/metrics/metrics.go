package metrics

import (
	"net/http"

	"github.com/avp365/go-exp/internal/services/metrics/handlers"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var MetricClient Client
var registry = prometheus.NewRegistry()

func init() {
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector())
}

type PromMetric interface {
	SetToPrometheus()
}

type Client struct {
	HandlersProcessingSend *handlers.ProcessingSend
}

func New() Client {
	client := Client{
		HandlersProcessingSend: handlers.NewProcessingSend(),
	}
	client.HandlersProcessingSend.Register(registry)

	return client
}

func (c Client) metrics() []PromMetric {
	return []PromMetric{
		c.HandlersProcessingSend,
	}
}

func (c Client) NewMetricsHandler() http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		for _, m := range c.metrics() {
			m.SetToPrometheus()
		}

		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(rsp, req)
	})
}
