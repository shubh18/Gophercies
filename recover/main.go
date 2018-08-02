package main

import (
	"log"
	"net/http"
	m "recover/Middleware"
)

var listenAndServeFunc = http.ListenAndServe

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", m.SourceCodeHandler)
	mux.HandleFunc("/panic/", m.PanicHandler)
	log.Fatal(listenAndServeFunc(":8000", m.Middleware(mux)))
}
