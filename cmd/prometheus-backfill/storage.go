package main

type Storage interface {
	Parser(b []byte) error
	Writer(b []byte) error
}

type StorageTSDB struct {
	MaxBatchSize uint64
}

type StorageInfluxDB struct {
	MaxBatchSize uint64
}

func NewStorageTSDB(batchSize uint64) (*StorageTSDB, error) {
	return &StorageTSDB{
		MaxBatchSize: batchSize,
	}, nil
}

func NewStorageInfluxDB(batchSize uint64) (*StorageInfluxDB, error) {
	return &StorageInfluxDB{
		MaxBatchSize: batchSize,
	}, nil
}

func (st *StorageTSDB) Parser(b []byte) error {

	return nil
}

func (st *StorageInfluxDB) Parser(b []byte) error {

	return nil
}
