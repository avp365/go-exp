package main

import (
	"fmt"
	"net/http"
)

func app(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "app\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusRequestTimeout)

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func health(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "ok")
}

func main() {

	http.HandleFunc("/", app)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/health", health)

	http.ListenAndServeTLS(":9200", ".infra/ssl/riva.local.crt", ".infra/ssl/riva.local.key", nil)

}
