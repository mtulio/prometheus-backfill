package storage

type Storage interface {
	Parser(b []byte) error
	Writer(p interface{}) error
	Close() error
}
