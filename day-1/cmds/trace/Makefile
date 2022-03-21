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