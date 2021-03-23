package backfill

type Backfill struct {
	options *Options
	scli    *StorageClient
	pcli    *ParserClient
}

func NewBackfill() *Backfill {
	return &Backfill{}
}

func NewBackfillWithOptions(opts *Options) (*Backfill, error) {
	bf := Backfill{options: opts}

	sc, err := NewStorageClientWithOptions(opts)
	if err != nil {
		return nil, err
	}
	bf.scli = sc

	pc, err := NewParserWithOptions(opts, sc)
	if err != nil {
		return nil, err
	}
	bf.pcli = pc
	return &bf, nil
}

func (b *Backfill) Close() {
	if b.scli.client != nil {
		b.scli.client.Close()
	}
	return
}

func (b *Backfill) StartParser() error {
	// buf, err := parseGzipFile(*b.options.ArgIn)
	// if err != nil {
	// 	return err
	// }
	// parseBuffer(buf, b.stg)

	b.pcli.parser.Parse()
	return nil
}
