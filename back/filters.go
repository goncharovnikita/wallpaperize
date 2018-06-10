package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func mFilter(method string, next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next.ServeHTTP(rw, r)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func contentLengthFilter(next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil || length == 0 {
			rw.WriteHeader(400)
			rw.Write([]byte("Empty content not allowed"))
		} else {
			next.ServeHTTP(rw, r)
		}
	}
}

func headersFilter(headers []string, next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		for _, v := range headers {
			hdr := r.Header.Get(v)
			if hdr == "" {
				rw.WriteHeader(400)
				rw.Write([]byte(fmt.Sprintf("header %s is not present", v)))
				return
			}
		}

		next.ServeHTTP(rw, r)
	}
}
