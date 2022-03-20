// based on:
// https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/blob/main/hello-app/main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var Payload MetricsPayload

func main() {
	// Generate metrics payload
	var err error
	Payload, err = GeneratePayload()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Metrics payload is: %+v\n", Payload)

	// Send metrics payload on startup
	err = PublishMetricsPayload(Payload)
	if err != nil {
		log.Fatal(err)
	}

	// Run periodic call-home ping in the background
	go PeriodicPing(Payload)

	// Register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/telemetry", showTelemetry)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "📊 Hello, telemetry!\n")
	fmt.Fprintf(w, "💻 Hostname: %s\n", host)
}

func showTelemetry(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	fmt.Fprintf(w, "🔎 Metrics Payload:\n")
	fmt.Fprintf(w, "%+v\n", Payload)
}
