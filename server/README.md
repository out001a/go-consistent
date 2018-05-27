# go-consistent HTTP Server

## Usage

### Build & Run

```bash
> go build server.go    # build it

> ./server -h           # get help
Usage of ./server:
  -port int
        The port to listen on. (default 9871)

> ./server              # run server
```

### Request

```bash
# add node
curl -v -XPOST http://127.0.0.1:9871/api/v1/consistent/ -d node1

# lookup node
curl -v -XGET http://127.0.0.1:9871/api/v1/consistent/key1

# remove node
curl -v -XDELETE http://127.0.0.1:9871/api/v1/consistent/ -d node1
```
