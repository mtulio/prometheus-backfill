package backfill

import (
	"errors"
	"fmt"

	"github.com/mtulio/prometheus-backfill/backfill/parser"
	"github.com/mtulio/prometheus-backfill/backfill/storage"
)

type ParserClient struct {
	options *Options
	scli    storage.Storage
	parser  parser.Parser
}

func NewParserWithOptions(opts *Options, sc *StorageClient) (*ParserClient, error) {
	pc := ParserClient{
		options: opts,
		scli:    sc.client,
	}
	if (pc.options.inType == "file") || (pc.options.inType == "directory") {
		var isDir bool = false
		if pc.options.inType == "directory" {
			isDir = true
		}
		fp, err := parser.NewFileParser(
			*(pc.options.ArgIn),
			*(pc.options.ArgInExt),
			isDir,
			pc.scli,
		)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error Creating Parser Client: %v", err))
		}
		pc.parser = fp
		return &pc, nil
	}

	return nil, errors.New(fmt.Sprintf("Error Creating Storage Client: Unknow Parser type %s", sc.options.inType))
}
