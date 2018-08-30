# play-wasm

My boilerplate for Go Web Assembly.

## Run

```shell
make build && make run
```

Then open your browser at [localhost:8080](http://localhost:8080).

## Description

The `make build` command executes a `docker build`. The Dockerfile uses the official Go 1.11 image
* to compile `play.go` into wasm
* to compile `serve/serve.go` into the web server that will serve the files.
