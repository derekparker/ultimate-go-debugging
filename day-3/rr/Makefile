.PHONY: set-perf-event-paranoid
set-perf-event-paranoid:
	echo 1 | sudo tee /proc/sys/kernel/perf_event_paranoid

.PHONY: check-perf-list
check-perf-list:
	perf list

.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/recordme main.go

.PHONY: rr-record
rr-record:
	rr record ./bin/recordme

.PHONY: dlv-replay
dlv-replay:
	dlv replay $(TRACE)