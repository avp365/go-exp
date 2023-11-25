package main

import (
	"net/http"
	"os"

	"github.com/avp365/go-exp/internal/services/metrics"

	formatter "github.com/fabienm/go-logrus-formatters"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
)

func init() {

	gelfFmt := formatter.NewGelf("go-exp")
	log.SetFormatter(gelfFmt)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	hook := graylog.NewGraylogHook("go-exp-graylog-3-3:12201", map[string]interface{}{})
	log.AddHook(hook)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Files received from server")

	log.WithFields(log.Fields{
		"requestID": "id99304050",
	}).Debug("Info User")
	metrics.MetricClient.HandlersProcessingSend.Add("send")

}
func main() {

	log.Println("metric server start ")

	metrics.MetricClient = metrics.New()
	http.Handle("/metrics", metrics.MetricClient.NewMetricsHandler())

	go func() {
		err := http.ListenAndServe(":7755", nil)
		if err != nil {
			log.Println(err)
		}
	}()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8100", nil))

}
