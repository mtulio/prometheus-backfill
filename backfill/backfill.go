package backfill

type Backfill struct {
	options *Options
	stg     *StorageClient
}

func NewBackfill() *Backfill {
	return &Backfill{}
}

func NewBackfillWithOptions(opts *Options) (*Backfill, error) {
	bf := Backfill{options: opts}

	stg, err := NewStorageClientWithOptions(opts)
	if err != nil {
		return nil, err
	}
	bf.stg = stg

	return &bf, nil
}

func (b *Backfill) Close() {
	if b.stg.client != nil {
		b.stg.client.Close()
	}
	return
}

func (b *Backfill) StartParser() error {
	buf, err := parseGzipFile(*b.options.ArgInTarget)
	if err != nil {
		return err
	}
	parseBuffer(buf, b.stg)
	return nil
}
