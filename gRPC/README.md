### The Math gRPC module

- ` math/math/math.proto` file holds the prototype buffer language code which structures our protocol buffer data.
- ` math/math/math.pb.go ` file is created using protocol buffer compiler with file types and API functions.
- ` math/server/server.go` file executes the server which accepts a `Request` message and sends a `Response` message.
- ` math/client/client.go` file executes the client which sends a `Request` messasge to server and prints the incoming response.

### SETTING UP FILES
Put math folder in `$GOPATH/src/`

#### Running the protocol buffer compiler
Protocol buffer compiler Creates \*.pb.go file from \*.proto file (if changes are done)
```
$ cd $GOPATH/src/math
$ protoc -I ./ --go_out=plugins=grpc:./ ./math/math.proto
```

### Running server
` $ ./server/server.go `

### Running client
` $ ./client/client.go `

