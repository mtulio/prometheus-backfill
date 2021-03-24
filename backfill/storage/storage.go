package storage

import "time"

type Storage interface {
	Parser(b []byte) error
	Writer(p interface{}) error
	NewPoint(
		name string,
		labels map[string]string,
		values map[string]interface{},
		timestamp time.Time,
	) interface{}
	Close() error
}
