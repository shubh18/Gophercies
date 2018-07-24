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
	mux.HandleFunc("/panic/", PanicHandler)
	mux.HandleFunc("/panic-after/", PanicAfterHandler)
	log.Fatal(http.ListenAndServe(":8000", m.Middleware(mux)))
}

//PanicHandler to handle panic function
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	PanicFunction()
}

//PanicAfterHandler to handle response after panic
func PanicAfterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Recover</h1>")
	PanicFunction()
}

//PanicFunction tp display panic
func PanicFunction() {
	panic("called")
}
