package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func maxVersionHandler(path string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		patch := 0
		minor := 0
		major := 0

		err := filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
			if info.IsDir() {
				return nil
			}

			rex := regexp.MustCompile("\\w+\\-\\w+\\-\\w\\.\\w\\.\\w")

			if !rex.MatchString(info.Name()) {
				return nil
			}

			nums := strings.Split(strings.Split(info.Name(), "-")[2], ".")
			ma, e1 := strconv.Atoi(nums[0])
			mi, e2 := strconv.Atoi(nums[1])
			pa, e3 := strconv.Atoi(nums[2])
			if e1 != nil || e2 != nil || e3 != nil {
				return nil
			}

			if ma > major {
				major = ma
				minor = mi
				patch = pa
			} else if ma == major {
				if mi > minor {
					minor = mi
					patch = pa
				} else if mi == minor {
					if pa > patch {
						patch = pa
					}
				}
			}
			return nil
		})

		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			return
		}

		rw.Write([]byte(fmt.Sprintf("%d.%d.%d", major, minor, patch)))
	}
}
