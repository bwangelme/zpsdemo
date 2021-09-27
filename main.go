package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/openzipkin/zipkin-go"
	zkhttpmw "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/openzipkin/zipkin-go/reporter"
	zkreporter "github.com/openzipkin/zipkin-go/reporter/kafka"
)

var (
	zkReporter reporter.Reporter
)

const (
	serviceName     = "zipkin_http_server"
	serviceEndpoint = "localhost:8080"
	kafkaAddr       = "localhost:2181"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
	return
}

func initMux() http.Handler {
	var err error

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Pong)

	address := []string{kafkaAddr}
	zkReporter, err = zkreporter.NewReporter(address)
	if err != nil {
		log.Fatalf("unable to connect to kafka: %+v\n", err)
	}
	endpoint, err := zipkin.NewEndpoint(serviceName, serviceEndpoint)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	tracer, err := zipkin.NewTracer(
		zkReporter, zipkin.WithTraceID128Bit(true),
		zipkin.WithLocalEndpoint(endpoint),
	)

	zkMiddleware := zkhttpmw.NewServerMiddleware(tracer)
	return zkMiddleware(mux)
}

func main() {
	mux := initMux()

	http.ListenAndServe(":8080", mux)
}
