bin/execme:
	go build -o bin/execme

.PHONY: dlv-exec
dlv-exec:
	dlv exec ./bin/execme

.PHONY: build-showing-inlining
build-showing-inlining:
	go build -gcflags=-m=2 -o bin/prog

.PHONY: search-dwarf-info
search-dwarf-info: bin/execme
	objdump --dwarf=info bin/execme | grep -C 2 Rectangle

bin/execmenooptimizations:
	go build -o bin/execmenooptimizations -gcflags="all=-N -l"

.PHONY: dlv-exec-no-optimizations
dlv-exec-no-optimizations:
	dlv exec ./bin/execmenooptimizations