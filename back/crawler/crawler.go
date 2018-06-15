package crawler

import "github.com/goncharovnikita/wallpaperize/app/api"

// RandomCrawler download random images from source
type RandomCrawler struct {
	saveDir  string
	master   api.Wallmaster
	maxSpace int64
}

// NewRandomCrawler create random crawler
func NewRandomCrawler(saveDir string, maxUsage int64, master api.Wallmaster) *RandomCrawler {
	return &RandomCrawler{
		saveDir:  saveDir,
		maxSpace: maxUsage,
		master:   master,
	}
}
