# go-consistent

Consistent Hash implementation in Go

## Usage

```go
import "github.com/out001a/go-consistent"

consist := consistent.NewConsistent()

consist.Add("node1")
consist.Add("node2")

v1, ok1 := consist.Lookup("key1")

consist.Remove("node1")

v2, ok2 := consist.Lookup("key1")
```

## PHP Version

https://github.com/out001a/consistent-hash/

## HTTP Server

It supplies a simple RESTful HTTP server, which can be found in the [server](/server) package.

## TODO

* perfect test cases
* use travis-ci
* gen godoc
* dockerize the server
* data persistence
* service metrics
* benchmark and optimize
* use namespace or database to support different types of nodes
