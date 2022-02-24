package gotool

import (
	"io/ioutil"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func DirRegListFiles(path string, findreg string) []*string {
	list := make([]string, 0)
	reg, _ := regexp.Compile(findreg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Errorln("open dir failed")
		return nil
	}

	for _, file := range files {
		filename := file.Name()
		if !file.IsDir() && reg.Match(filename) {
			list = append(list, file.Name())
		}
	}

	return &list
}

func DirRegListDirs(path string, findreg string) []*string {
	list := make([]string, 0)
	reg, _ := regexp.Compile(findreg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Errorln("open dir failed")
		return nil
	}

	for _, file := range files {
		filename := file.Name()
		if file.IsDir() && reg.Match(filename) {
			list = append(list, file.Name())
		}
	}

	return &list
}
