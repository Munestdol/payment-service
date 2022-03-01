FROM golang:1.17-alpine3.15 AS builder

COPY . /payment-service/
WORKDIR /payment-service/

RUN go mod download
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /payment-service/bin/app .
COPY --from=builder /payment-service/configs configs/
COPY --from=builder /payment-service/migrations migrations/

EXPOSE 80 50080

CMD ["./app"]