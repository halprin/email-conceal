GOOS = darwin
GOARCH = amd64

compile: compileForwarder compileManager

compileForLinux: GOOS = linux
compileForLinux: GOARCH = amd64
compileForLinux: compile

compileForwarder:
	cd ./src/ && \
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="-w -s" -o forwarder -v ./cmd/forwarder/

compileManager:
	cd ./src/ && \
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="-w -s" -o manager -v ./cmd/manager/

test:
	go test ./...

runLocally:
	docker-compose up

devDeploy: compileForLinux
	cd ./iac/environments/dev/ && \
	terraform apply -auto-approve
