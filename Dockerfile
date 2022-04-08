FROM golang:1.18 AS builder
COPY pingbox.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 go build pingbox.go

FROM busybox:latest
COPY --from=builder /build/pingbox /app/
WORKDIR /app
CMD ["/app/pingbox"]

