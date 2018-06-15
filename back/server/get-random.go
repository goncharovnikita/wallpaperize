package server

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/goncharovnikita/wallpaperize/back/utils"
)

func (s *Server) handleGetRandom() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rn := rand.New(rand.NewSource(time.Now().UnixNano()))
		fnames, err := utils.GetFilenamesFromDir(s.randomPath)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("error occured"))
			return
		}

		result := make([]string, 0, 10)
		for i := 0; i < 10; i++ {
			index := rn.Intn(len(fnames))
			result = append(result, fnames[index])
			fnames = append(fnames[:index], fnames[index+1:]...)
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
