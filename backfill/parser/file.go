package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mtulio/prometheus-backfill/backfill/storage"
)

type FileParser struct {
	path    string
	isDir   bool
	fileExt string
	files   []string
	scli    storage.Storage
}

func NewFileParser(path, fileExt string, isDir bool, scli storage.Storage) (*FileParser, error) {
	fp := FileParser{
		path:    path,
		fileExt: fileExt,
		isDir:   isDir,
		scli:    scli,
	}
	if isDir {
		fmatch := fmt.Sprintf("%s/*.%s", fp.path, fp.fileExt)
		matches, err := filepath.Glob(fmatch)
		if err != nil {
			return nil, errors.New("Error Directory found but no match *.json.gz")
		}
		fp.files = matches
	} else {
		fp.files = []string{fp.path}
	}

	return &fp, nil
}

func (p *FileParser) Parse() error {

	log.Printf("Total files to be processed: %v\n", len(p.files))
	var cnt uint = 1
	var tt int = len(p.files)
	for _, f := range p.files {
		log.Printf("Processing file %d/%d: %s\n", cnt, tt, f)

		buf, err := p.Reader(f)
		if err != nil {
			//return err
			log.Println("Error reader %v\n", err)
			continue
		}
		//parseBuffer(buf, p.scli)
		err = p.Parser(buf)
		if err != nil {
			log.Println("Error parsing %v\n", err)
		}
		cnt += 1
	}
	return nil
}

func (p *FileParser) Close() error {

	return nil
}

// Reader reads a file from disk, unpacking it if needed
func (p *FileParser) Reader(f string) ([]byte, error) {

	var fReader io.ReadCloser

	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// Only the first 512 bytes are used to sniff the content type.
	fHeader := make([]byte, 512)
	_, err = fd.Read(fHeader)
	if err != nil {
		log.Println(err)
	}
	fd.Seek(0, 0)
	ftype := http.DetectContentType(fHeader)

	log.Printf("Detected file type: %s\n", ftype)

	if ftype == "application/x-gzip" {
		fReader, err = gzip.NewReader(fd)
		if err != nil {
			return nil, err
		}
		defer fReader.Close()
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, fReader)

	return buf.Bytes(), nil
}

// Reader reads a file from disk, unpacking it if needed
func (p *FileParser) Parser(buf []byte) error {

	var (
		promJson  PrometheusResponse
		ttMetrics uint64
		ttPoints  uint64
	)

	err := json.Unmarshal(buf, &promJson)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v\n", err)
	}

	log.Printf("JSON parsed: status=%s, ResultType=%s, #Labels=%d\n",
		promJson.Status, promJson.Data.ResultType,
		len(promJson.Data.Result))
	log.Println("Processing data points...")

	for idxM := range promJson.Data.Result {
		ttMetrics += 1
		for idxP := range promJson.Data.Result[idxM].Values {
			ttPoints += 1
			ts := time.Unix(int64(promJson.Data.Result[idxM].Values[idxP][0].(float64)), 0)
			value, _ := strconv.ParseFloat(promJson.Data.Result[idxM].Values[idxP][1].(string), 64)
			name := promJson.Data.Result[idxM].Metric["__name__"]

			p.scli.Writer(p.scli.NewPoint(
				name,
				promJson.Data.Result[idxM].Metric,
				map[string]interface{}{
					"value": value,
				},
				ts,
			))
			if (ttPoints % 100000) == 0 {
				log.Printf("Processing data points...%d\n", ttPoints)
			}
		}
	}
	log.Printf("Processed %d data points\n", ttPoints)
	return nil
}
