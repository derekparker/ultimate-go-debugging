.PHONY: build-crash
build-crash:
	go build -gcflags="all=-N -l" -o bin/crash crash.go

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash

.PHONY: set-core-pattern
set-core-pattern:
	echo core | sudo tee /proc/sys/kernel/core_pattern

.PHONY: dlv-core
dlv-core:
	dlv core ./bin/crash core.$(PID)

.PHONY: gcore
gcore:
	gcore $(PID)

.PHONY: dlv-exec
dlv-exec:
	dlv exec ./bin/crash