package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhaogaolong/graphql-metrics/graphql"
)

//go:generate ./patch/graphql_patch.sh

func main() {

	http.HandleFunc("/", graphql.GraphIQLHandler)
	http.HandleFunc("/graphql", graphql.GraphQLHandler)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":3000", nil))
}
