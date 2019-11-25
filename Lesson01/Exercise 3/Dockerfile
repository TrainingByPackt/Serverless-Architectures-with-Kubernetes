FROM golang:1.12.5-alpine3.9 AS builder
ADD . .
RUN go build *.go

FROM alpine:3.9
COPY --from=builder /go/function ./function
RUN chmod +x ./function
ENTRYPOINT ["./function"]