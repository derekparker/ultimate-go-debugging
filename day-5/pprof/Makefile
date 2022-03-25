.PHONY: build
build:
	go build -o bin/prog

.PHONY: run
run:
	./bin/prog

.PHONY: get-heap-profile
get-heap-profile:
	curl http://localhost:6060/debug/pprof/heap > heap.pprof

.PHONY: go-tool-pprof-heap
go-tool-pprof-heap:
	go tool pprof heap.pprof

.PHONY: run-cpu-profile
run-cpu-profile:
	./bin/prog -cpuprofile cpu.pprof

.PHONY: view-cpu-profile
view-cpu-profile:
	go tool pprof ./bin/prog ./cpu.pprof

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