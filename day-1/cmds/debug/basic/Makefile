.PHONY: run
run:
	go run last-character.go

.PHONY: install-delve
install-delve:
	go install github.com/go-delve/delve/cmd/dlv@latest

.PHONY: debug
debug:
	dlv debug

.PHONY: debug-with-arg
debug-with-arg:
	dlv debug -- "foobar"