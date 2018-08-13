package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

func main() {

	res, err := http.Get("https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS")
	if err != nil {
		log.Fatalf("error fetching: %v", err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error parsing the response body: %v", err)
	}

	document := js.Global().Get("document")
	p := document.Call("getElementById", "target")
	p.Set("innerHTML", string(b))
}
