FROM golang:1.12.5-alpine3.9 AS builder
RUN apk add --no-cache git

ADD main.go /go/src/gold-price-to-mysql/main.go
WORKDIR /go/src/gold-price-to-mysql/
RUN go get -v
RUN go build .

FROM alpine:3.9
COPY --from=builder /go/src/gold-price-to-mysql/gold-price-to-mysql ./gold-price-to-mysql
RUN chmod +x ./gold-price-to-mysql
ENTRYPOINT ["./gold-price-to-mysql"]