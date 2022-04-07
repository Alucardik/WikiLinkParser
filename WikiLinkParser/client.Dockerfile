FROM golang:1.18

WORKDIR /usr/local/wiki-parser_server

COPY config config
COPY error_utils error_utils
COPY limiter limiter
COPY queue_info queue_info
COPY proto proto
COPY client client
COPY worker worker
COPY server server
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod tidy
RUN go build .

CMD ./WikiLinkParser --mode=client
