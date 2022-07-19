bin/traceme:
	go build -gcflags="all=-N -l" -o bin/traceme

.PHONY: strace
strace:
	strace ./bin/traceme

.PHONY: trace
trace:
	dlv trace ^main.callme
