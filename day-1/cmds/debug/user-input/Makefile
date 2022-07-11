PHONY: run
run:
	go run main.go

.PHONY: debug
debug:
	dlv debug

.PHONY: debug-with-tty
debug-with-tty:
	dlv debug --tty=$(TTY)

.PHONY: install-ptyme
install-ptyme:
	git clone https://github.com/derekparker/ptyme.git && cd ptyme && cargo build --release

.PHONY: run-ptyme
run-ptyme: install-ptyme
	cd ptyme && cargo run

.PHONY: debug-headless
debug-headless:
	dlv debug --headless --listen :9090 --api-version 2

.PHONY: connect-to-headless-delve
connect-to-headless-delve:
	dlv connect :9090