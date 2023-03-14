bin/tracme:
	go build -gcflags="all=-N -l" -o bin/traceme main.go

.PHONY: strace
strace:
	strace ./bin/traceme

.PHONY: trace
trace:
	dlv trace main.go ^main.callme

.PHONY: trace-with-stack
trace-with-stack:
	dlv trace main.go ^main.callme --stack 2

.PHONY: trace-with-ebpf
trace-with-ebpf:
	sudo env "PATH=$$PATH" dlv trace main.go callme --ebpf