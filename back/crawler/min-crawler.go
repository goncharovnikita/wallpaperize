package crawler

import (
	"image/jpeg"
	"log"
	"strings"
	"time"
	// Import jpeg codecs
	_ "image/jpeg"
	"os"
)

// MinCrawler minimifies images and save to distinct directory
type MinCrawler struct {
	dirs           []string
	minFilePostfix string
	minDirPostfix  string
}

// NewMinCrawler creates new MinCrawler
func NewMinCrawler(dirs []string) *MinCrawler {
	return &MinCrawler{
		dirs:           dirs,
		minFilePostfix: "-min",
		minDirPostfix:  "_min",
	}
}

// Crawl imlementation
func (c *MinCrawler) Crawl() {
	log.Println("min crawler starts")
	for _, v := range c.dirs {
		os.Mkdir(v+c.minDirPostfix, 0777)
	}
	for {
		log.Println("min crawler crawling")
		for _, v := range c.dirs {
			log.Println("min crawler enter ", v)
			c.crawl(v)
			log.Println("min crawler ends at ", v)
		}
		time.Sleep(time.Second * 5)
	}
}

func (c *MinCrawler) crawl(dir string) {
	minDir := dir + c.minDirPostfix

	filesToMin := make([]string, 0)
	filesToRemove := make([]string, 0)

	fnames, err := getFilenamesFromDir(dir)

	if err != nil {
		log.Println(err)
		return
	}

	minfnames, err := getFilenamesFromDir(minDir)

	if err != nil {
		log.Println(err)
		return
	}

	for _, fname := range fnames {
		finded := false
		for _, minname := range minfnames {
			rname := strings.Replace(minname, c.minFilePostfix, "", 1)
			if rname == fname {
				finded = true
				break
			}
		}

		if !finded {
			filesToMin = append(filesToMin, fname)
		}
	}

	for _, minname := range minfnames {
		finded := false
		rname := strings.Replace(minname, c.minFilePostfix, "", 1)
		for _, fname := range fnames {
			if rname == fname {
				finded = true
				break
			}
		}

		if !finded {
			filesToRemove = append(filesToRemove, minname)
		}
	}

	for _, v := range filesToRemove {
		fname := minDir + "/" + v
		log.Println("removing ", v)
		os.Remove(fname)
	}

	for _, v := range filesToMin {
		err = c.min(v, dir, minDir)
		if err != nil {
			log.Println("cannot minimify image")
		}
	}
}

// fname - target file name, e.g. test.jpg
// fpath - target file path, e.g. /images/
// distpath - destination file path, e.g. /images_min/
func (c *MinCrawler) min(fname, fpath, distpath string) error {
	log.Println("minimifying ", fname)
	minname := strings.Replace(fname, ".jpg", "", 1) + c.minFilePostfix + ".jpg"
	minfile, err := os.OpenFile(distpath+"/"+minname, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	defer minfile.Close()

	file, err := os.OpenFile(fpath+"/"+fname, os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return err
	}

	err = jpeg.Encode(minfile, img, &jpeg.Options{
		Quality: 50,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
