package main

import (
	"net/http"
)

func serve(path string) {
	http.HandleFunc(
		"/add/build",
		mFilter("POST",
			contentLengthFilter(
				headersFilter(
					[]string{VERSION_HEADER},
					addBuild(path),
				),
			),
		),
	)

	http.HandleFunc(
		"/get/maxversion",
		corsHeader(maxVersionHandler(path)),
	)
}
