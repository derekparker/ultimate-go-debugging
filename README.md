# Ultimate Go Debugging Course

This repo holds the example code and course information for the Ultimate Go Debugging course.

This course will teach you various techniques that can be used to debug a Go program.

Cloning this repo:

```
git clone https://github.com/derekparker/ultimate-go-debugging.git --branch march2023
```

## Getting Started

### Installing required tools

### Go

First, you should have the latest version of Go. You can get the latest version from https://go.dev/dl/.
Once you download the tarball follow the instructions on https://go.dev/doc/install to install Go on your system.
To verify a correct Go install run the following command:

```
go version
```

You should see output similar to:

```
go version go1.17.6 linux/amd64
```

### Delve

Second, once you have Go installed you must install Delve. You can install the latest version of Delve via:

```
go install github.com/go-delve/delve/cmd/dlv@latest
```

You can verify Delve has been correctly installed by running the following command:

```
dlv help
```

**NOTE**: If you're on macOS you may require a few more steps to get fully up and running, so please refer to the [documentation](https://github.com/go-delve/delve/tree/master/Documentation/installation#macos-considerations).

The full install documentation is [here](https://github.com/go-delve/delve/tree/master/Documentation/installation) if needed.

### Vagrant

For remote debugging we will be using Vagrant.

For instructions on installing, [click here](https://developer.hashicorp.com/vagrant/docs/installation).

### Container runtime

You must have a container runtime, ideally one that is CLI compatible with Docker.

* For instructions on installing Docker [click here](https://docs.docker.com/get-docker/).
* For instructions on installing Podman [click here](https://podman.io/getting-started/installation).

### Kind

Lastly, you will need a way to spin up a development Kubernetes cluster. For the workshop we will be using Kind.

For installation instructions, [click here](https://kind.sigs.k8s.io/docs/user/quick-start/#installation).

## What we will cover

### Day 1

* Basics of using Debugging
* Intro into CLI usage
* Navigating your program
* Inspecting / changing program state

### Day 2 (pt 1)

* Advanced program navigation
* Tracing your program
* How to debug core dumps

### Day 3 (pt 2)

* Advanced program navigation
* Tracing your program
* How to debug core dumps

### Day 4

* Scripting Delve
* Remote debugging
* How to use JSON-RPC API
* Record and replay debugging


### Day 5 (pt 1)

* Debugging in containers
* Debugging in Kubernetes

### Day 6 (pt 2)

* Debugging in containers
* Debugging in Kubernetes

### Day 7 (pt 1)

* Debugging through performance analysis
* How to use pprof
* How to use perf
* Deep dive into Delve internals
* Deep dive into Go internals

### Day 8 (pt 2)

* Debugging through performance analysis
* How to use pprof
* How to use perf
* Deep dive into Delve internals
* Deep dive into Go internals
* Wrapup