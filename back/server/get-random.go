package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/goncharovnikita/wallpaperize/back/utils"
)

func (s *Server) handleGetRandom() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fnames, err := utils.GetFilenamesFromDir(s.randomPath)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("error occured"))
			return
		}

		data, err := json.Marshal(fnames)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("error occured"))
			return
		}

		rw.Write(data)

	}
}
