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

.PHONY: run-basic-image
run-basic-image:
	docker run -it --detach -p 8080:8080 --rm $(BASIC_IMG)