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

.PHONY: objdump-debug-info
objdump-debug-info:
	objdump -g ./bin/trace > debug-info.txt

.PHONY: go-objdump
go-objdump:
	go tool objdump -gnu -s "main.callme" -S ./bin/trace > go-objdump-disassemble.txt
