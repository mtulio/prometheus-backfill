# Builder image
FROM golang:1.16-alpine3.13 AS builder

RUN apk update && apk add --no-cache git

ENV APP=prometheus-backfill
WORKDIR $GOPATH/src/github.com/mtulio/${APP}

COPY . .
RUN go get -d -v ./...
RUN mkdir -p /go/bin \
    && go build -o /go/bin/${APP} github.com/mtulio/${APP}/cmd/${APP}

# Main image
FROM alpine:3.13

ENV APP=prometheus-backfill

COPY --from=builder /go/bin/${APP} /${APP}
CMD ["/prometheus-backfill"]