package storage

import (
	"github.com/mtulio/prometheus-backfill/pkg/backfill/storage"
)

type Storage interface {
	Parser(b []byte) error
	Writer(b []byte) error
	Close()
}

// type StorageTSDB struct {
// 	MaxBatchSize uint64
// }

// type StorageInfluxDB struct {
// 	MaxBatchSize uint64
// }

func NewStorageClient() *Storage {

	return storage.NewStorageInfluxDB{}
}
