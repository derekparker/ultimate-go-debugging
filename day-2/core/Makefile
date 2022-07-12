.PHONY: build-crash
build-crash:
	go build -o bin/crash crash.go

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash