package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/goncharovnikita/wallpaperize/back/models"
)

func (s *Server) handleGetRandom() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var limit int

		limitStr := r.URL.Query().Get("limit")
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

func (s *Server) handleGetRandomImage() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		images, err := s.imagesGetter.GetImages(1)
		if err != nil {
			s.logger.Printf("error get image: %v\n", err)

			rw.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(rw).Encode(models.ResponseError{
				Error: "could not get image",
			}); err != nil {
				s.logger.Printf("failed to write response: %v\n", err)
			}

			return
		}

		if len(images) != 1 {
			rw.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(rw).Encode(models.ResponseError{
				Error: "no images in database",
			}); err != nil {
				s.logger.Printf("failed to write response: %v\n", err)
			}

			return
		}

		image := images[0]
		imageURL := image.URLs.Full

		switch r.URL.Query().Get("size") {
		case "small":
			imageURL = image.URLs.Small
		case "raw":
			imageURL = image.URLs.RAW
		}

		http.Redirect(rw, r, imageURL, http.StatusFound)
	}
}
