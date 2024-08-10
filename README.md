# Custom-VM-Container-GoLang

This Go project demonstrates how to create a simple container-like environment using Linux namespaces. It shows how to spawn a new process with isolated namespaces for UTS (hostname and domain name), PID (process IDs), and filesystem mounts.

## Overview

The project contains a single Go program that creates a new process in an isolated environment. The new process runs in a separate UTS, PID, and mount namespace, which is similar to how containers work.

### Components

1. **`main.go`**: The main file containing the logic for spawning a new process with isolated namespaces.

2. **Namespaces Used**:
   - **UTS Namespace**: Isolates the hostname and domain name.
   - **PID Namespace**: Isolates the process IDs.
   - **Mount Namespace**: Isolates the filesystem mounts.

## Usage

### Building the Project

To build the project, run:

```sh
go build -o container main.go

