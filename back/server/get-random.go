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
		var limit int

		limitStr := chi.URLParam(r, "limit")
		if limitStr == "" {
			limit = 20
		} else {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				s.logger.Printf("got request with bad limit: %s\n", limitStr)

				rw.WriteHeader(http.StatusBadRequest)
				if err := json.NewEncoder(rw).Encode(models.ResponseError{
					Error: "limit should be integer",
				}); err != nil {
					s.logger.Printf("failed to write response: %v\n", err)
				}

				return
			}

			limit = limitInt
		}

		images, err := s.imagesGetter.GetImages(limit)
		if err != nil {
			s.logger.Printf("error get images: %v\n", err)

			rw.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(rw).Encode(models.ResponseError{
				Error: "could not get images",
			}); err != nil {
				s.logger.Printf("failed to write response: %v\n", err)
			}

			return
		}

		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(images); err != nil {
			s.logger.Printf("failed to write response: %v\n", err)
		}
	}
}
