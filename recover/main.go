package main

import (
	"fmt"
	"log"
	"net/http"
	m "recover/Middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", m.SourceCodeHandler)
	mux.HandleFunc("/panic/", m.PanicHandler)
	mux.HandleFunc("/panic-after/", m.PanicAfterHandler)
	log.Fatal(http.ListenAndServe(":8000", m.Middleware(mux)))
}
