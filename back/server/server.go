package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/goncharovnikita/wallpaperize/back/models"
	"github.com/sirupsen/logrus"
)

const (
	VERSION_HEADER = "BUILD_VERSION"
)

type imagesGetter interface {
	GetImages(limit int) ([]*models.ResponseImage, error)
}

// Server type
type Server struct {
	buildPath    string
	randomPath   string
	debug        bool
	imagesGetter imagesGetter
	logger       *logrus.Logger
}

// NewServer creates new server
func NewServer(bp, rp string, imagesGetter imagesGetter, debug bool) *Server {
	return &Server{
		buildPath:    bp,
		randomPath:   rp,
		imagesGetter: imagesGetter,
		debug:        debug,
	}
}

func (s *Server) Listen() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

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

	return r
}

func getLogger(logger *logrus.Logger, r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"path": r.URL.Path,
	})
}
