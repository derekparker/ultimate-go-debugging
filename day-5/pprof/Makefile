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