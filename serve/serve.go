package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func file(name string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", time.Now().Format(time.RFC3339), req.Method, req.RequestURI)
		if strings.HasSuffix(name, ".wasm") {
			rw.Header().Set("Content-Type", "application/wasm")
		}
		http.ServeFile(rw, req, name)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", file("./serve/wasm_exec.html"))
	mux.Handle("/wasm_exec.js", file("./serve/wasm_exec.js"))
	mux.Handle("/play.wasm", file("./serve/play.wasm"))

	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
