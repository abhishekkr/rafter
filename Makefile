PACKAGE_NAME := github.com/abhishekkr/rafter

# compress golang binary size: https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
build: deps
	ls -al
	rm -rf artifacts

	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-server $(PACKAGE_NAME)/cmd/server
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-server.osx $(PACKAGE_NAME)/cmd/server
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-server.exe $(PACKAGE_NAME)/cmd/server

	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-client $(PACKAGE_NAME)/cmd/client
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-client.osx $(PACKAGE_NAME)/cmd/client
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows go build -ldflags="-s -w" -mod=mod -a -installsuffix cgo -o artifacts/rafter-client.exe $(PACKAGE_NAME)/cmd/client

deps:
	go mod tidy
