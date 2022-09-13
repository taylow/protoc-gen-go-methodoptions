# protoc-gen-go-methodoptions

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

This is a protoc plugin that generates Go code for the `google.protobuf.MethodOptions` message type. This allows you to write your own custom options for your gRPC services, and provides a few ways to access them (constants, request methods, and service methods).

## Getting Started <a name = "getting_started"></a>

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Installing


```bash
go install github.com/taylow/protoc-gen-go-methodoptions@latest
```

### Generate

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-methodoptions_out=. --go-methodoptions_opt=paths=source_relative example/service.proto
```

## Usage <a name = "usage"></a>

This system was developed while working at [Comnoco](https://github.com/comnoco) and is used to allow us to add custom options to our gRPC services. We use this to define our authorisation roles at the proto level, allowing both the frontend and backend to use the same authorisation rules. A full example of this can be found in the [example](example) directory.
