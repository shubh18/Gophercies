package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	link "recover/Hyperlink"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

//SourceCodeHandler function to handle source code of given file
func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	filepath := r.FormValue("path")
	lineNo := r.FormValue("line")
	line, err := strconv.Atoi(lineNo)
	if err != nil {
		line = -1
	}
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, buffer.String())
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.TabWidth(2), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	formatter.Format(w, style, iterator)
}

//Middleware to recover panics from program
func Middleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, link.CreateLinks(string(stack)))
			}
		}()
		handler.ServeHTTP(w, r)
	}
}
