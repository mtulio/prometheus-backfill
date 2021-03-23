package backfill

import (
	"errors"
	"fmt"

	"github.com/mtulio/prometheus-backfill/backfill/storage"
)

type StorageClient struct {
	options *Options
	client  storage.Storage
}

func NewStorageClientWithOptions(opts *Options) (*StorageClient, error) {
	sc := StorageClient{options: opts}

	if sc.options.outStorageType == "influxdb" {
		stg, err := storage.NewStorageInfluxDB(
			sc.options.outStorageAddr,
			sc.options.outStorageDb,
			sc.options.outStorageAuth,
			uint(opts.backendBufferSize),
		)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error Creating Storage Client: %v", err))
		}
		sc.client = stg
		return &sc, nil
	}

	return nil, errors.New(fmt.Sprintf("Error Creating Storage Client: Unknow storage type %s", sc.options.outStorageType))
}
