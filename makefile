
build_local_server:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/server/cred-server-linux github.com/yurchenkosv/credential_storage/cmd/credentialsServer

build_local_client:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/client/cred-client-linux github.com/yurchenkosv/credential_storage/cmd/credentialsClient
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/client/cred-client-darwin github.com/yurchenkosv/credential_storage/cmd/credentialsClient
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/client/cred-client-windows github.com/yurchenkosv/credential_storage/cmd/credentialsClient

certs:
	cd hack && ./gen_certs.sh && cd ..

run_server: certs build_local_server
	chmod +x bin/server/cred-server-linux
	docker run -d -e POSTGRES_DB=credentials -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:14
	sleep 5
	hack/run_server.sh

test:
	go test ./...
