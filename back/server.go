package main

import (
	"net/http"
)

func serve(path string) {
	http.HandleFunc(
		"/add/build",
		mFilter("PUT", contentLengthFilter(addBuild(path))),
	)
}
