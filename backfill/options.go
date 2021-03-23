package backfill

import (
	"errors"
	"fmt"
	"strings"
)

type Options struct {
	backendBufferSize uint64
	inputBufferSize   uint64
	parserBufferSize  uint64

	// param ArgOutTarget: stgType=address=db=user=pass
	outStorageType string
	outStorageAddr string
	outStorageDb   string
	outStorageAuth string

	ArgInTarget  *string
	ArgOutTarget *string
	ArgBatchSize *uint64
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

func (o *Options) Parse() error {
	stgParams := strings.Split(*o.ArgOutTarget, "=")
	if len(stgParams) != 5 {
		return errors.New("Error on -i argument. Expect: stgType=address=db=user=pass")
	}
	o.outStorageType = stgParams[0]
	o.outStorageAddr = stgParams[1]
	o.outStorageDb = stgParams[2]
	o.outStorageAuth = fmt.Sprintf("%s:%s", stgParams[3], stgParams[4])

	o.backendBufferSize = uint64(*o.ArgBatchSize)

	return nil
}
