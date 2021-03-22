package storage

func NewStorageTSDB(batchSize uint64) (*StorageTSDB, error) {
	return &StorageTSDB{
		MaxBatchSize: batchSize,
	}, nil
}

func (st *StorageTSDB) Parser(b []byte) error {

	return nil
}
