trie:
	git clone git@github.com:derekparker/trie.git

.PHONY: run-benchmarks
run-benchmarks:
	GO111MODULE=off go test -v -bench=".*" ./trie

.PHONY: run-benchmarks-with-profiling
run-benchmarks-with-profiling:
	GO111MODULE=off go test -v -cpuprofile cpu.prof -memprofile mem.prof -bench="." ./trie

.PHONY: view-cpu-profile
view-cpu-profile:
	go tool pprof cpu.prof