FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/drprado2/transaction-manager/
WORKDIR /go/src/github.com/drprado2/transaction-manager
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/transaction-manager /go/src/github.com/drprado2/transaction-manager/cmd/transaction-manager/transaction-manager.go

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/drprado2/transaction-manager/build/transaction-manager /usr/bin/transaction-manager
COPY --from=builder /go/src/github.com/drprado2/transaction-manager/pkg/storage/postgres-db/migrations /usr/migrations
RUN ls -la /usr/migrations
EXPOSE 9000 9000
ENTRYPOINT ["/usr/bin/transaction-manager"]