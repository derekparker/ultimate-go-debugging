.PHONY: build
build:
	go build -o bin/myclient

.PHONY: run
run:
	./bin/myclient localhost:9090 | jq .

.PHONY: run-testprog-with-dlv
run-testprog-with-dlv:
	cd testprog && dlv debug --continue --headless --api-version=2 --accept-multiclient --listen :9090

.PHONY: kill-dlv-server
kill-dlv-server:
	pkill dlv

.PHONY: run-with-args
run-with-args:
	./bin/myclient localhost:9090 main.setupHandlers.func1 | jq .

.PHONY: curl-app
curl-app:
	curl localhost:8081/foo/bar