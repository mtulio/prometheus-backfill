package parser

type Parser interface {
	Parse() error
	Close() error
}
