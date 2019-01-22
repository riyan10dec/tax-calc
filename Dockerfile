# Build Go
FROM golang:alpine AS builder

# Install Git for go get
RUN set -eux; \
    apk add --no-cache --virtual git
ENV GO_WORKDIR $GOPATH/src/github.com/riyan10dec/tax-calc/
RUN mkdir -p src/github.com/riyan10dec/tax-calc/
ADD . ${GO_WORKDIR}
WORKDIR ${GO_WORKDIR} 
RUN go get -v
# RUN ./go-bindata -ignore=\\.go -pkg=schema -o=schema/bindata.go schema/...
RUN go install 

# Minimize Docker Size
FROM alpine:latest
RUN set -eux; \
    apk add --no-cache --virtual ca-certificates
COPY --from=builder /go/bin/tax-calc .
COPY --from=builder /go/src/github.com/riyan10dec/tax-calc/config.toml .
RUN apk add --no-cache tzdata
CMD ["./tax-calc"]
EXPOSE 8080