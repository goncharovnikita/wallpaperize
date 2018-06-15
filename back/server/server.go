package server

import (
	"log"
	"net/http"
)

const (
	VERSION_HEADER = "BUILD_VERSION"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Server type
type Server struct {
	buildPath  string
	randomPath string
}

// NewServer creates new server
func NewServer(bp, rp string) *Server {
	return &Server{
		buildPath:  bp,
		randomPath: rp,
	}
}

// Serve bootstraps server
func (s *Server) Serve() {
	s.serve()
}

func (s *Server) serve() {
	http.HandleFunc(
		"/add/build",
		mFilter("POST",
			contentLengthFilter(
				headersFilter(
					[]string{VERSION_HEADER},
					addBuild(s.buildPath),
				),
			),
		),
	)

	http.HandleFunc(
		"/get/maxversion",
		corsHeader(maxVersionHandler(s.buildPath)),
	)

	http.HandleFunc(
		"/get/random",
		corsHeader(
			addHeadersFilter(
				map[string]string{"Content-Type": "application/json"},
				s.handleGetRandom())),
	)
}
