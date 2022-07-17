FROM golang:1.16-alpine AS builder
WORKDIR /build
COPY . /build
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 go build -ldflags="-w -s"

FROM alpine AS runner
WORKDIR /build
COPY --from=builder /build/cloud-tool .
EXPOSE 80
CMD ["./cloud-tool"]