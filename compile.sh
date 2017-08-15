docker run --rm --net="host" -v "$PWD":/usr/src/myapp -v "$PWD/gopath":/go -w /usr/src/myapp golang:1.8.3-alpine3.6 go build server.go
