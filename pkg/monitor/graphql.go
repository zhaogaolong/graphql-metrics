package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(GraphqlMetrics)
	prometheus.MustRegister(GraphqlErrorMetrics)
}

var (
	GraphqlMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "graphql_query_mutation_method_total",
			Help:        "count the graphql query and mutation method",
			ConstLabels: prometheus.Labels{"service": "medusa"},
		},
		[]string{"method", "type"},
	)
	GraphqlErrorMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "graphql_method_error_total",
			Help:        "count the graphql api errors",
			ConstLabels: prometheus.Labels{"service": "medusa"},
		},
		[]string{"method", "type"},
	)
)
