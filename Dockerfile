FROM golang:1.19 AS Builder
WORKDIR /go/src/github.com/yurchenkosv/credential_storage/
COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o cred_storage_server cmd/credentialsServer/server.go

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=Builder /go/src/github.com/yurchenkosv/credential_storage/cred_storage_server ./
CMD ["./cred_storage_server"]