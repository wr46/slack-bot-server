FROM golang:1.19-bullseye as builder

RUN apt-get install ca-certificates && update-ca-certificates

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY . .

RUN go build -a -ldflags "-s -w" -o main ./cmd

FROM debian:bookworm-slim

RUN groupadd -r app && useradd --no-log-init -r -g app app

WORKDIR /app
COPY --from=builder --chown=app /src/main /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app

CMD ["./main"]
