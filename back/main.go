package main

import (
	"net/http"
	"os"
)

const (
	VERSION_HEADER = "BUILD_VERSION"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	path := os.Getenv("BUILDS_PATH")
	if path == "" {
		path = "uploads"
		os.Mkdir("uploads", 0777)
	}

	serve(path)
	println("app listen on :: ", port)

	http.ListenAndServe(":"+port, nil)
}
