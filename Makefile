compile: compileForwarder compileManager

compileForwarder:
	cd ./src/ && \
	go build -ldflags="-w -s" -o forwarder -v ./cmd/forwarder/

compileManager:
	cd ./src/ && \
	go build -ldflags="-w -s" -o manager -v ./cmd/manager/

test:
	go test ./...

runLocally:
	docker-compose up

devDeploy: compile
	cd ./iac/environments/dev/ && \
	terraform apply -auto-approve
