package storage

import (
	"context"
	"net/url"
	"sort"
	"sync"
	"time"

	"log"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/storage/remote"
)

type StorageRemote struct {
	maxBatchSize   uint
	maxClientConn  uint
	clientWriter   remote.WriteClient
	bufferWriter   chan *prompb.WriteRequest
	bufferWriterWg sync.WaitGroup
	queueWriter    *prompb.WriteRequest
}

func NewStorageRemote(name, address string) (*StorageRemote, error) {
	var batchSz uint = 20
	var maxWriters uint = 2

	u2, err := url.Parse(address)
	if err != nil {
		log.Fatal(err)
	}
	u1 := u2.(config_util.URL)

	cfg := remote.ClientConfig{URL: u2.(*config_util.URL)}
	wc, err := remote.NewWriteClient(name, &cfg)
	if err != nil {
		log.Println(err)
	}
	qw := &prompb.WriteRequest{
		Timeseries: make([]prompb.TimeSeries, 0, 0),
	}
	sr := StorageRemote{
		maxBatchSize:  batchSz,
		maxClientConn: maxWriters,
		clientWriter:  wc,
		bufferWriter:  make(chan *prompb.WriteRequest, batchSz),
		queueWriter:   qw,
	}

	// starting writers
	for i := uint(0); i < sr.maxClientConn; i++ {
		sr.bufferWriterWg.Add(1)
		go func() {
			defer sr.bufferWriterWg.Done()
			for r := range sr.bufferWriter {
				if err := sr.Commit(r); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}

	return &sr, nil
}

// func NewRemoteStorage() {
// 	var (
// 		remoteStorage = remote.NewStorage(log.With(logger, "component", "remote"), prometheus.DefaultRegisterer, localStorage.StartTime, cfg.localStoragePath, time.Duration(cfg.RemoteFlushDeadline), scraper)
// 		fanoutStorage = storage.NewFanout(logger, localStorage, remoteStorage)
// 	)
// }

func (s *StorageRemote) Writer(metrics interface{}) error {
	stream := metrics.([]*model.SampleStream)

	totalSamples := uint(0)
	for _, st := range stream {
		samples := make([]prompb.Sample, 0, len(st.Values))
		for _, v := range st.Values {
			samples = append(samples, prompb.Sample{
				Value:     float64(v.Value),
				Timestamp: int64(v.Timestamp),
			})
			totalSamples++
		}
		ts := prompb.TimeSeries{
			Labels:  metricToLabelProtos(st.Metric),
			Samples: samples,
		}
		s.queueWriter.Timeseries = append(s.queueWriter.Timeseries, ts)
		if totalSamples > s.maxBatchSize {
			log.Printf("Sending batch of %d samples", totalSamples)
			totalSamples = 0
			s.bufferWriter <- s.queueWriter
			s.queueWriter = &prompb.WriteRequest{Timeseries: make([]prompb.TimeSeries, 0, 0)}
		}
	}

	log.Printf("Sending batch of %d samples", totalSamples)
	s.bufferWriter <- s.queueWriter
	return nil
}

func metricToLabelProtos(metric model.Metric) []prompb.Label {
	labels := make([]prompb.Label, 0, len(metric))
	for k, v := range metric {
		labels = append(labels, prompb.Label{
			Name:  string(k),
			Value: string(v),
		})
	}
	sort.Slice(labels, func(i int, j int) bool {
		return labels[i].Name < labels[j].Name
	})
	return labels
}

func (s *StorageRemote) Commit(req *prompb.WriteRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60)
	defer cancel()

	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	compressed := snappy.Encode(nil, data)
	s.clientWriter.Store(ctx, compressed)
	return nil
}

func (s *StorageRemote) Parser(b []byte) error {

	return nil
}

func (s *StorageRemote) Close() error {

	return nil
}

func (s *StorageRemote) NewPoint(
	name string,
	labels map[string]string,
	values map[string]interface{},
	timestamp time.Time,
) interface{} {
	return nil
}

func (s *StorageRemote) NewStream(
	name string,
	labels map[string]string,
	values map[string]interface{},
	timestamp time.Time,
) interface{} {
	return nil
}
