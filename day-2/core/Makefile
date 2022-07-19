.PHONY: build-crash
build-crash:
	go build -gcflags="all=-N -l" -o bin/crash crash.go

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash