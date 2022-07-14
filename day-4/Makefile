.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/prog ./main.go

.PHONY: copy-dlv
copy-dlv:
	cp $$(which dlv) ./bin

.PHONY: start-vm
start-vm:
	vagrant up

.PHONY: destroy-vm
destroy-vm:
	vagrant destroy

.PHONY: ssh-to-vm
ssh-to-vm:
	vagrant ssh

.PHONY: port-forward
port-forward:
	vagrant ssh -- -NL 4040:localhost:4040

.PHONY: connect-to-remote
connect-to-remote:
	dlv connect :4040

##############################
## Container Image commands ##
##############################

BASIC_IMG := buggy-basic
SCRATCH_IMG := buggy-scratch
DELVE_IMG := buggy-with-delve

# Build basic container image.
.PHONY: build-image
build-image:
	docker build --pull --rm -f build/Dockerfile-basic -t $(BASIC_IMG):latest .
