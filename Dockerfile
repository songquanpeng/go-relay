FROM golang AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -X 'go-relay/common.Version=$(cat VERSION)'" -o go-relay

FROM scratch

COPY --from=builder /build/go-relay /
EXPOSE 6872
WORKDIR /app
ENTRYPOINT ["/go-relay"]