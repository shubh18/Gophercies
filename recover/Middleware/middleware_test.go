package middleware

import "net/http"

func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		panic("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}
