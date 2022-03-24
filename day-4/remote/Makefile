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