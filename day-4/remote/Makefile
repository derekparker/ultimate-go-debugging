.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/prog

.PHONY: copy-dlv
copy-dlv:
	cp $$(which dlv) ./bin

.PHONY: start-vm
start-vm:
	vagrant up

.PHONY: ssh-to-vm
ssh-to-vm:
	vagrant ssh

.PHONY: port-forward
port-forward:
	vagrant ssh -- -NL 4040:localhost:4040

.PHONY: connect-to-remote
connect-to-remote:
	dlv connect :4040

.PHONY: start-dap-server
start-dap-server:
	dlv dap --listen localhost:4040

.PHONY: dlv-exec-headless
dlv-exec-headless:
	dlv exec --headless --accept-multiclient --listen :4040 /vagrant/bin/prog