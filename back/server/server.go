package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
	debug      bool
}

// NewServer creates new server
func NewServer(bp, rp string, debug bool) *Server {
	return &Server{
		buildPath:  bp,
		randomPath: rp,
		debug:      debug,
	}
}

// Serve bootstraps server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serve(w, r)
}

func (s *Server) serve(rw http.ResponseWriter, request *http.Request) {
	log.Println("server serving")
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	allowedOrigins := make([]string, 0)

	if s.debug {
		allowedOrigins = []string{"*"}
		r.Use(middleware.Logger)
	} else {
		allowedOrigins = []string{"https://wallpaperize.goncharovnikita.com"}
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/add/build", addBuild(s.buildPath))

	r.Get(
		"/get/maxversion",
		maxVersionHandler(s.buildPath),
	)

	r.Get(
		"/get/random",
		addHeadersFilter(
			map[string]string{"Content-Type": "application/json"},
			s.handleGetRandom()),
	)

	r.Get(
		"/get/links",
		addHeadersFilter(
			map[string]string{"Content-Type": "application/json"},
			s.handleGetDownloadLinks()),
	)

	r.ServeHTTP(rw, request)
}
