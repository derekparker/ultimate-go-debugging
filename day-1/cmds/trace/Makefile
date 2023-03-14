bin/tracme:
	go build -gcflags="all=-N -l" -o bin/traceme main.go

.PHONY: strace
strace:
	strace ./bin/traceme