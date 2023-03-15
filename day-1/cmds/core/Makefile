bin/crash:
	go build -o bin/crash crash.go

.PHONY: run-crash
run-crash: bin/crash
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash

.PHONY: ulimit-unlimited
ulimit-unlimited:
	ulimit -c unlimited

.PHONY: set-core-pattern
set-core-pattern:
	echo core | sudo tee /proc/sys/kernel/core_pattern

.PHONY: dlv-core
dlv-core:
	dlv core ./bin/crash core.$(PID)