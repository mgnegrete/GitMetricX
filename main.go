package main

import (
	"fmt"
	"net/http"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "# HELP gitmetricx_commits_total Total number of commits")
	fmt.Fprintln(w, "# TYPE gitmetricx_commits_total counter")
	fmt.Fprintln(w, "gitmetricx_commits_total 42")
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
