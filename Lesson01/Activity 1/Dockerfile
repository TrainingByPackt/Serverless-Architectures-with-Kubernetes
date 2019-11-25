FROM golang:1.12.5-alpine3.9 AS builder
ADD . .
RUN go build *.go

FROM alpine:3.9
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
COPY --from=builder /go/function ./bikes
RUN chmod +x ./bikes
ENTRYPOINT ["./bikes"]