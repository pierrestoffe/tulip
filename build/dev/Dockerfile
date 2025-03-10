FROM golang:1.24.1-alpine

WORKDIR /app/

RUN go install github.com/cespare/reflex@latest

COPY ./go.mod ./
RUN go mod download

COPY ./build/dev/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]
