package storage

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2api "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type StorageInfluxDB struct {
	MaxBatchSize uint
	client       influxdb2.Client
	clientWAPI   influxdb2api.WriteAPI
}

func NewStorageInfluxDB(addr, db, auth string, batchSize uint) (*StorageInfluxDB, error) {
	si := StorageInfluxDB{MaxBatchSize: batchSize}
	si.client = influxdb2.NewClientWithOptions(
		addr, auth,
		influxdb2.DefaultOptions().SetBatchSize(si.MaxBatchSize),
	)
	si.clientWAPI = si.client.WriteAPI("default", db)
	return &si, nil
}

func (si *StorageInfluxDB) Parser(b []byte) error {

	return nil
}

func (si *StorageInfluxDB) Close() error {
	si.clientWAPI.Flush()
	si.client.Close()
	return nil
}

func (si *StorageInfluxDB) Writer(p interface{}) error {
	si.clientWAPI.WritePoint(p.(*write.Point))
	return nil
}

func (si *StorageInfluxDB) NewPoint(
	name string,
	labels map[string]string,
	values map[string]interface{},
	timestamp time.Time,
) interface{} {
	return influxdb2.NewPoint(
		name,
		labels,
		values,
		timestamp)
}
