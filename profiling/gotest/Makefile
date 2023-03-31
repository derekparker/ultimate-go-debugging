trie:
	git clone git@github.com:derekparker/trie.git

.PHONY: run-benchmarks
run-benchmarks:
	cd trie && go test -v -bench ".*"

.PHONY: run-benchmarks-with-profiling
run-benchmarks-with-profiling:
	cd trie && go test -v -cpuprofile cpu.prof -memprofile mem.prof -bench ".*"

.PHONY: view-cpu-profile
view-cpu-profile:
	go tool pprof trie/cpu.prof

.PHONY: view-memory-profile
view-memory-profile:
	go tool pprof trie/mem.prof
