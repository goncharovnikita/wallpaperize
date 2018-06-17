package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/github"
)

// DownloadLinks type
type DownloadLinks struct {
	Mac   string `json:"mac"`
	Linux string `json:"linux"`
}

func (s Server) handleGetDownloadLinks() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		client := github.NewClient(nil)
		releases, _, err := client.Repositories.ListReleases(context.Background(), "goncharovnikita", "wallpaperize", nil)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("an error occurred"))
			return
		}

		if len(releases) < 1 {
			rw.WriteHeader(400)
			rw.Write([]byte("repo dont have releases"))
			return
		}

		release := releases[0]

		result := DownloadLinks{}
		for _, v := range release.Assets {
			name := *(v.Name)
			if strings.HasSuffix(name, ".dmg") {
				result.Mac = *(v.BrowserDownloadURL)
				continue
			}

			if strings.HasSuffix(name, ".AppImage") {
				result.Linux = *(v.BrowserDownloadURL)
				continue
			}
		}

		err = json.NewEncoder(rw).Encode(result)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("error occurred"))
			return
		}
	}
}
