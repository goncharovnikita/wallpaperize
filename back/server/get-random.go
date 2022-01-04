package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/goncharovnikita/wallpaperize/back/models"
)

func (s *Server) handleGetRandom() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log := getLogger(s.logger, r)
		var limit int

		limitStr := chi.URLParam(r, "limit")
		if limitStr == "" {
			limit = 20
		} else {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				log.Warnf("got request with bad limit: %s", limit)

				rw.WriteHeader(http.StatusBadRequest)
				if err := json.NewEncoder(rw).Encode(models.ResponseError{
					Error: "limit should be integer",
				}); err != nil {
					log.Errorf("failed to write response: %w", err)
				}

				return
			}

			limit = limitInt
		}

		images, err := s.imagesGetter.GetImages(limit)
		if err != nil {
			s.logger.Errorf("error get images: %w", err)

			rw.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(rw).Encode(models.ResponseError{
				Error: "could not get images",
			}); err != nil {
				log.Errorf("failed to write response: %w", err)
			}
		}

		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(images); err != nil {
			log.Errorf("failed to write response: %w", err)
		}
	}
}
