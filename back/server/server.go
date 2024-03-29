package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/goncharovnikita/wallpaperize/back/internal/models"
	pubmodels "github.com/goncharovnikita/wallpaperize/back/models"
)

const (
	VERSION_HEADER = "BUILD_VERSION"
)

type imagesGetter interface {
	GetImages(limit int) ([]*pubmodels.UnsplashImage, error)
}

// Server type
type Server struct {
	buildPath    string
	randomPath   string
	debug        bool
	imagesGetter imagesGetter
	logger       *log.Logger
}

// NewServer creates new server
func NewServer(
	bp, rp string,
	imagesGetter imagesGetter,
	logger *log.Logger,
	debug bool,
) *Server {
	return &Server{
		buildPath:    bp,
		randomPath:   rp,
		imagesGetter: imagesGetter,
		debug:        debug,
		logger:       logger,
	}
}

func (s *Server) Listen() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	var allowedOrigins []string

	if s.debug {
		allowedOrigins = []string{"*"}
	} else {
		allowedOrigins = []string{
			"https://wallpaperize.goncharovnikita.com",
			"https://goncharovnikita.com",
		}
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
			s.handleGetRandom(),
		),
	)

	r.Get(
		"/get/links",
		addHeadersFilter(
			map[string]string{"Content-Type": "application/json"},
			s.handleGetDownloadLinks(),
		),
	)

	r.Get("/random/image.jpg", s.handleGetRandomImage())

	return r
}

type logger interface {
	Printf(msg string, args ...interface{})
}

func write500(rw http.ResponseWriter, r *http.Request, log logger, err error) {
	rw.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(rw).Encode(models.ResponseError{
		Error: err.Error(),
	}); err != nil {
		log.Printf("failed to write response: %v\n", err)
	}
}
