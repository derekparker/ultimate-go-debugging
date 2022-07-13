.PHONY: build
build:
	go build -o bin/myclient

.PHONY: run
run:
	./bin/myclient localhost:9090 | jq .

.PHONY: run-testprog-with-dlv
run-testprog-with-dlv:
	dlv debug --headless --api-version=2 --accept-multiclient --listen :9090

.PHONY: kill-dlv-server
kill-dlv-server:
	pkill dlv