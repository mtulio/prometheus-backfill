package backfill

type Options struct {
	backendBufferSize uint64
	inputBufferSize   uint64
	parserBufferSize  uint64
}

func (o *Options) BackendBufferSize() uint64 {
	return o.backendBufferSize
}

func (o *Options) SetBackendBufferSize(value uint64) *Options {
	o.backendBufferSize = value
	return o
}

func (o *Options) InputBufferSize() uint64 {
	return o.inputBufferSize
}

func (o *Options) SetInputBufferSize(value uint64) *Options {
	o.inputBufferSize = value
	return o
}

func (o *Options) ParserBufferSize() uint64 {
	return o.parserBufferSize
}

func (o *Options) SetParserBufferSize(value uint64) *Options {
	o.parserBufferSize = value
	return o
}

func (o *Options) SetBuffersSize(vBack, vInput, vParser uint64) *Options {
	o.SetBackendBufferSize(vBack)
	o.SetInputBufferSize(vInput)
	o.SetParserBufferSize(vParser)
	return o
}
