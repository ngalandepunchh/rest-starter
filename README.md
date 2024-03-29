# rest-starter [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/hunterpunchh/rest-starter) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterpunchh/rest-starter?style=flat-square)](https://goreportcard.com/report/github.com/hunterpunchh/rest-starter) [![Release](https://img.shields.io/github/release/hunterpunchh/rest-starter.svg?style=flat-square)](https://github.com/hunterpunchh/rest-starter/releases/latest)

## Prerequisites

* grpc-starter: See https://github.com/hunterpunchh/grpc-starter for setup

## Setup

* Build the container `make build`
* Compose the containers `make up`

## Example

```bash
$ curl localhost:8080/status

# Expected response
{
    "status":"ok"
}
```

```bash
$ curl localhost:8080/healthz

# Expected response
{
    "status":"ok",
    "grpcStarterStatus":"ok"
}
```
