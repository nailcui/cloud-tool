FROM golang:alpine

WORKDIR /build
COPY cloud-tool .

EXPOSE 80
CMD ["/build/cloud-tool"]
