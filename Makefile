
build:
	docker build -t play-wasm .

run: build
	docker run --rm -p 8080:8080 play-wasm
