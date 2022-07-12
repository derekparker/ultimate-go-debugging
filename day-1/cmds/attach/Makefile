.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/test-attach-program

.PHONY: run
run:
	./bin/test-attach-program

.PHONY: curl-app
curl-app:
	curl localhost:8081/foo/bar

.PHONY: attach-debugger
attach-debugger:
	dlv attach $$(pgrep -f test-attach-program)