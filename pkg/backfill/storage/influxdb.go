package storage

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2api "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type StorageInfluxDB struct {
	MaxBatchSize uint64
	client       influxdb2.Client
	clientWAPI   influxdb2api.WriteAPI
}

func NewStorageInfluxDB(addr, db, auth string, batchSz uint) (*StorageInfluxDB, error) {
	var si StorageInfluxDB
	si.client = influxdb2.NewClientWithOptions(
		addr, auth,
		influxdb2.DefaultOptions().SetBatchSize(batchSz),
	)
	si.clientWAPI = si.client.WriteAPI("default", db)
	return &si, nil
}

func (si *StorageInfluxDB) Parser(b []byte) error {

	return nil
}

func (si *StorageInfluxDB) Writer(b []byte) error {

	return nil
}

func (si *StorageInfluxDB) Close() {
	si.clientWAPI.Flush()
	si.client.Close()
	return
}

func (si *StorageInfluxDB) WritePoint(p *write.Point) error {
	si.clientWAPI.WritePoint(p)
	return nil
}
