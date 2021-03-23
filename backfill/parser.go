package backfill

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func unzip(source string) ([]byte, error) {
	reader, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	archive, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer archive.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, archive)
	// fmt.Println("buf read from gz file :", buf)
	return buf.Bytes(), nil
}

func parseGzipFile(path string) ([]byte, error) {
	buf, err := unzip(path)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

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

func parseBuffer(buf []byte, stg *StorageClient) error {
	// Unmarshal using a generic interface
	//writeAPI

	fmt.Println("Parsing...")
	var genericJSON PrometheusResponse
	err := json.Unmarshal(buf, &genericJSON)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Println("Parsing...2")
	fmt.Println(genericJSON.Status)
	fmt.Println(genericJSON.Data.ResultType)

	fmt.Println("Parsing...3")
	var ttMetrics uint64
	var ttPoints uint64
	for idxM := range genericJSON.Data.Result {
		ttMetrics += 1
		for idxP := range genericJSON.Data.Result[idxM].Values {
			ttPoints += 1
			ts := time.Unix(int64(genericJSON.Data.Result[idxM].Values[idxP][0].(float64)), 0)
			value, _ := strconv.ParseFloat(genericJSON.Data.Result[idxM].Values[idxP][1].(string), 64)
			name := genericJSON.Data.Result[idxM].Metric["__name__"]

			p := influxdb2.NewPoint(
				name,
				genericJSON.Data.Result[idxM].Metric,
				map[string]interface{}{
					"value": value,
				},
				ts)

			//writeAPI.WritePoint(p)
			//client.WritePoint(p)
			stg.client.Writer(p)
		}
		if ttPoints%10000 == 0 {
			fmt.Printf("Total metrics: %d\n", ttMetrics)
			fmt.Printf("Total dpoints: %d\n", ttPoints)
		}
	}
	fmt.Printf("Total metrics: %d\n", ttMetrics)
	fmt.Printf("Total dpoints: %d\n", ttPoints)
	return nil
}
