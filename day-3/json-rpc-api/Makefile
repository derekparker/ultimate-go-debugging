.PHONY: build
build:
	go build -o bin/myclient

.PHONY: run
run:
	./bin/myclient localhost:9999 | jq .

.PHONY: run-testprog-with-dlv
run-testprog-with-dlv:
	cd ./testprog && dlv debug --continue --api-version=2 --headless --accept-multiclient --listen :9999

.PHONY: kill-dlv-server
kill-dlv-server:
	pkill dlv