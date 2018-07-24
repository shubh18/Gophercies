# Exercise #15: Recover Middleware using Chroma library

## Exercise details
This is my version of [Gophercise 15](https://github.com/gophercises/recover_chroma).This exercise is used to navigate to any source file in the panic stacktrace in order to make it easier to debug issues when panics get invoked.

Above repository contains following functions:

#### 1. An HTTP handler that will render source files in the browser

It is used to render pages if a specific path prefix (like `/debug/`) is used in the path.

#### 2. Add syntax highlighting to the source file rendering

Used the [chroma](https://github.com/alecthomas/chroma) package to add highlighting of source code.

#### 3. Parse the stack trace & creating links

Parse the stack trace and create links so that it can be used to identify errors.
For this step you can to use [url.Values](https://golang.org/pkg/net/url/#Values) to encode the path to the source file.

#### 4. Add line highlighting

Chroma supports [line highlighting](https://github.com/alecthomas/chroma#the-html-formatter) and stack trace has the line number where the panic occurred.So use that line number to highlight particular line to trace the error.
