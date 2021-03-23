
BIN ?= ./bin
PROJECT ?=  github.com/mtulio/prometheus-backfill
APP ?= prometheus-backfill

build:
	go build -o $(BIN)/$(APP) $(PROJECT)/cmd/$(APP)
