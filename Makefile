.PHONY: test.unit
test.unit:
	echo "=> Running Unit Tests"
	go test ./...

.PHONY: build
build:
	echo "=> Building the server"
	go build

.PHONY: build.docker
build.docker:
	echo "=> Building the docker image"
	docker build -t iboware/location-history .

.PHONY: generate
generate:
	oapi-codegen -package server -generate server api/oapi.yaml > pkg/server/location-server.gen.go
	oapi-codegen -package server -generate types api/oapi.yaml > pkg/model/location.gen.go