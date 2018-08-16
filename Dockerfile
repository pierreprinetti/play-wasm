FROM golang:1.11rc1

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/play-wasm
COPY . ./
RUN GOOS=js GOARCH=wasm go build -o ./serve/play.wasm ./play.go
RUN GOOS=linux go build -o ./app ./serve/serve.go

EXPOSE 8080

CMD ./app
