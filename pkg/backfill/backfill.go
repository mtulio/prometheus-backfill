package backfill

type Backfill struct {
	BackendBufferSize uint64
	InputBufferSize   uint64
	ParserBufferSize  uint64
}

func NewBackfill() *Backfill {
	return &Backfill{}
}

func NewBackfillWithOptions(opt *Options) *Backfill {
	return &Backfill{}
}
