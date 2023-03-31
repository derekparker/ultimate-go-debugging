.PHONY: build
build:
	go build -o bin/prog

.PHONY: run
run:
	./bin/prog

.PHONY: get-heap-profile
get-heap-profile:
	curl http://localhost:6060/debug/pprof/heap > heap.pprof

.PHONY: get-allocation-profile
get-allocation-profile:
	curl http://localhost:6060/debug/pprof/allocation > allocation.pprof

.PHONY: get-cpu-profile
get-cpu-profile:
	curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.pprof

.PHONY: go-tool-pprof-allocation
go-tool-pprof-allocation:
	go tool pprof allocation.pprof

.PHONY: go-tool-pprof-heap
go-tool-pprof-heap:
	go tool pprof heap.pprof

.PHONY: go-tool-pprof-cpu
go-tool-pprof-cpu:
	go tool pprof cpu.pprof

.PHONY: run-perf-stat
run-perf-stat:
	perf stat ./bin/prog

.PHONY: run-perf-record
run-perf-record:
	perf record ./bin/prog

.PHONY: run-perf-report
run-perf-report:
	perf report

.PHONY: run-perf-annotate
run-perf-annotate:
	perf annotate