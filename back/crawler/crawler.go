package crawler

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/goncharovnikita/wallpaperize/app/api"
)

// RandomCrawler download random images from source
type RandomCrawler struct {
	saveDir   string
	getter    api.ImageGetter
	maxSpace  int64
	spaceLeft int64
}

// NewRandomCrawler create random crawler
func NewRandomCrawler(saveDir string, maxUsage int64, getter api.ImageGetter) *RandomCrawler {
	return &RandomCrawler{
		saveDir:  saveDir,
		maxSpace: maxUsage,
		getter:   getter,
	}
}

// Crawl implementation
func (r *RandomCrawler) Crawl() {
	size, err := getDirSize(r.saveDir)
	if err != nil {
		log.Println("crawler cannot save images", err)
		return
	}

	log.Println(r.maxSpace, size)

	if (r.maxSpace - size) < (10 * (1024 * 1024)) {
		log.Println("crawler took all free space")
		return
	}

	(*r).spaceLeft = r.maxSpace - size

	r.crawl()
}

func (r *RandomCrawler) crawl() {
	log.Println("## crawling ##")
	img, err := r.getter.GetImage()
	if err != nil {
		log.Println(err)
		time.Sleep(time.Minute * 15)
		r.crawl()
		return
	}

	h256 := sha256.New()
	h256.Write(img)
	hsh := fmt.Sprintf("%x", h256.Sum(nil))
	name := r.saveDir + "/" + hsh + ".jpg"
	size := len(img)

	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		time.Sleep(time.Second * 1)
		r.crawl()
		return
	}

	defer file.Close()
	io.Copy(file, bytes.NewReader(img))

	(*r).spaceLeft = (r.spaceLeft - int64(size))
	if r.spaceLeft > (10 * 1024 * 1024) {
		r.crawl()
	} else {
		log.Println("crawler fill all space")
	}
}
