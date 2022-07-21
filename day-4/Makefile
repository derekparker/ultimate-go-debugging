.PHONY: build
build:
	go build -gcflags="all=-N -l" -o bin/prog

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

.PHONY: start-headless-server
start-headless-server:
	dlv exec --check-go-version=false --headless --api-version=2 --listen localhost:4040 /vagrant/bin/prog

.PHONY: port-forward
port-forward:
	vagrant ssh -- -NL 4040:localhost:4040

.PHONY: connect-to-remote
connect-to-remote:
	dlv connect localhost:4040

##############################
## Container Image commands ##
##############################

BASIC_IMG := buggy-basic
SCRATCH_IMG := buggy-scratch
DELVE_IMG := buggy-with-delve

.PHONY: curl-app
curl-app:
	curl localhost:8080/foo/bar

# Build basic container image.
.PHONY: build-image
build-image:
	docker build --rm -f build/Dockerfile-basic -t $(BASIC_IMG):latest .

.PHONY: run-basic-image
run-basic-image:
	docker run -it --detach -p 8080:8080 --rm $(BASIC_IMG)

# Copy dlv binary into basic container.
.PHONY: copy-dlv-to-container
copy-dlv-to-container:
	docker cp $$(which dlv) $$(docker ps -aqf "ancestor=$(BASIC_IMG)"):/dlv

# Exec dlv within basic container.
.PHONY: exec-dlv-basic-container
exec-dlv-basic-container:
	docker exec -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv attach 1

# Exec dlv within basic container.
.PHONY: exec-dlv-basic-container-privileged
exec-dlv-basic-container-privileged:
	docker exec --privileged -it $$(docker ps -aqf "ancestor=$(BASIC_IMG)") /dlv attach 1

.PHONY: change-ptrace-yama
change-ptrace-yama:
	echo "0" | sudo tee /proc/sys/kernel/yama/ptrace_scope

# Run basic image with ptrace SYS_CAP.
.PHONY: run-basic-image-with-ptrace
run-basic-image-with-ptrace:
	docker run -it --detach --rm -p 8080:8080 --cap-add=SYS_PTRACE $(BASIC_IMG)
