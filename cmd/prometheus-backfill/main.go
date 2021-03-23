package main

import (
	"flag"

	"github.com/mtulio/prometheus-backfill/backfill"
)

var (
	DefaultBackendBufferSize uint64 = 10000
	DefaultInputBufferSize   uint64 = 10000
	DefaultParserBufferSize  uint64 = 10000
)

func main() {

	opts := new(backfill.Options)

	opts.ArgInTarget = flag.String("i", "/path/to/directory", "input file or directory")
	opts.ArgOutTarget = flag.String("o", "influxdb=address=db=username=pass", "output address")
	opts.ArgBatchSize = flag.Uint64("b", DefaultParserBufferSize, "Batch size")

	flag.Parse()
	opts.Parse()

	bf, err := backfill.NewBackfillWithOptions(opts)
	if err != nil {
		panic(err)
	}
	defer bf.Close()

	err = bf.StartParser()
	if err != nil {
		panic(err)
		return
	}
}
