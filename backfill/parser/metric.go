package parser

import "time"

type MetricValue struct {
	Timestap time.Time
	Value    string
}

type PrometheusResultMetric struct {
	Metric map[string]string `json:"metric"`
	Values [][]interface{}   `json:"values"`
}

type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string                   `json:"resultType"`
		Result     []PrometheusResultMetric `json:"result"`
	} `json:"data"`
}
