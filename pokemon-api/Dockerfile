FROM golang:1.20-bookworm as builder

WORKDIR /app

ADD . /app

RUN go build -o bin/pokemon-api .

FROM debian:bookworm-slim

COPY --from=builder /app/bin/pokemon-api .

EXPOSE 8080

CMD ["./pokemon-api"]

