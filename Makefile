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
	cd ./src/ && \
	go test ./...

runLocally:
	docker-compose up

devDeploy: compileForLinux
	cd ./iac/environments/dev/ && \
	terraform init && \
	terraform apply

prodDeploy: compileForLinux
	cd ./iac/environments/prod/ && \
	terraform init && \
	terraform apply -auto-approve -var 'domain=$(DOMAIN)'
