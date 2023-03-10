bin/execme:
	go build -o bin/execme main.go

.PHONY: dlv-exec
dlv-exec:
	dlv exec ./bin/execme

.PHONY: build-showing-inlining
build-showing-inlining:
	go build -gcflags=-m=2 -o bin/prog main.go