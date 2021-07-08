package backfill

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	backendBufferSize uint64
	inputBufferSize   uint64
	parserBufferSize  uint64

	// parsed from param ArgOutTarget: stgType=address=db=user=pass
	outType        string
	outStorageAddr string
	outStorageDb   string
	outStorageAuth string

	// parsed from param ArgIn: /path/to/file_or_dir
	inType string

	ArgIn        *string
	ArgInExt     *string
	ArgOut       *string
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

// ParseArgIn parse input argument and check if it exists on FS.
func (o *Options) ParseArgIn() error {
	// Fill Input flow argument (ArgIn)
	inParams := strings.Split(*o.ArgIn, "=")
	if len(inParams) > 1 {
		return errors.New("Error on -i argument. Expect format: /path/to/file")
	}

	if _, err := os.Stat(*o.ArgIn); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("Does not exists: %s", *o.ArgIn))
	}

	fi, err := os.Stat(*o.ArgIn)
	if err != nil {
		return errors.New(fmt.Sprintf("Stat error: %s", err))
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		o.inType = "directory"
	case mode.IsRegular():
		o.inType = "file"
	}

	return nil
}

// ParseArgOut parse output argument and split into backend conn params.
func (o *Options) ParseArgOut() error {
	stgParams := strings.Split(*o.ArgOut, "=")
	if len(stgParams) != 5 {
		return errors.New("Error on -o argument. Expect format: stgType=address=db=user=pass")
	}
	o.outType = stgParams[0]
	o.outStorageAddr = stgParams[1]
	o.outStorageDb = stgParams[2]
	o.outStorageAuth = fmt.Sprintf("%s:%s", stgParams[3], stgParams[4])

	o.backendBufferSize = uint64(*o.ArgBatchSize)

	return nil
}

func (o *Options) Parse() error {

	err := o.ParseArgIn()
	if err != nil {
		return err
	}

	err = o.ParseArgOut()
	if err != nil {
		return err
	}

	return nil
}
