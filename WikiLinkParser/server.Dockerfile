FROM golang:1.18

WORKDIR /usr/local/wiki-parser_server

COPY config config
COPY limiter limiter
COPY queue_info queue_info
COPY error_utils error_utils
COPY client client
COPY server server
COPY worker worker
COPY proto proto
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod tidy
RUN go build .

CMD ["bash", "-c", "while ! curl -s rabbitmq:15672 > /dev/null; do echo waiting for rabbitmq; sleep 3; done; ./WikiLinkParser"]
