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
	Mac     string `json:"mac"`
	Linux   string `json:"linux"`
	Windows string `json:"windows"`
}

func (s Server) handleGetDownloadLinks() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		client := github.NewClient(nil)
		release, _, err := client.Repositories.GetLatestRelease(context.Background(), "goncharovnikita", "wallpaperize")
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			rw.Write([]byte("an error occurred"))
			return
		}

		if release == nil {
			rw.WriteHeader(400)
			rw.Write([]byte("repo dont have releases"))
			return
		}

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

			if strings.HasSuffix(name, ".exe") {
				result.Windows = *(v.BrowserDownloadURL)
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
