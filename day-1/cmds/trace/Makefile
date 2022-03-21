bin/traceme:
	go build -gcflags="all=-N -l" -o bin/traceme

.PHONY: strace
strace:
	strace ./bin/traceme

.PHONY: trace
trace:
	dlv trace ^main.callme

.PHONY: trace-with-stack
trace-with-stack:
	dlv trace ^main.callme --stack 10

.PHONY: trace-with-ebpf
trace-with-ebpf:
	sudo env "PATH=$$PATH" dlv trace callme --ebpf
