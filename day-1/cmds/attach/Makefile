.PHONY: build
build:
	go build -o bin/test-attach-program

.PHONY: run
run:
	./bin/test-attach-program

.PHONY: curl-app
curl-app:
	curl localhost:8082/foo/bar

.PHONY: attach-debugger
attach-debugger:
	dlv attach $$(pgrep -f test-attach-program)