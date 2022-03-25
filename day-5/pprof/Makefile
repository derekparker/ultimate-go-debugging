.PHONY: build
build:
	go build -o bin/prog

.PHONY: run
run:
	./bin/prog

.PHONY: get-heap-profile
get-heap-profile:
	curl http://localhost:6060/debug/pprof/heap > heap.pprof

.PHONY: get-cpu-profile
get-cpu-profile:
	curl http://localhost:6060/debug/pprof/cpu > cpu.pprof

.PHONY: go-tool-pprof-heap
go-tool-pprof-heap:
	go tool pprof heap.pprof

.PHONY: go-tool-pprof-cpu
go-tool-pprof-cpu:
	go tool pprof cpu.pprof