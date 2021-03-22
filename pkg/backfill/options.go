package backfill

import (
	"fmt"
	"strings"
)

type Options struct {
	backendBufferSize uint64
	inputBufferSize   uint64
	parserBufferSize  uint64

	// param: stgType=address=db=user=pass
	outStorageType string
	outStorageAddr string
	outStorageDb   string
	outStorageUser string
	outStoragePass string

	FInTarget  *string
	FOutTarget *string
	FBatchSz   *uint64
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

func (o *Options) FlagParse() error {
	stgParams := strings.Split(*o.FInTarget, "=")
	if len(stgParams) != 5 {
		return fmt.Errorf("Error on -i argument. Expect: stgType=address=db=user=pass")
	}
	o.outStorageType = stgParams[0]
	o.outStorageAddr = stgParams[1]
	o.outStorageDb = stgParams[2]
	o.outStorageUser = stgParams[3]
	o.outStoragePass = stgParams[4]

	o.backendBufferSize = uint64(*o.FBatchSz)

	return nil
}
