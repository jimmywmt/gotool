package gotool

import (
	"io/ioutil"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func DirRegListFiles(path string, findReg string) []*string {
	list := make([]*string, 0)
	reg, _ := regexp.Compile(findReg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Errorln("open dir failed")
		return nil
	}

	reg2, _ := regexp.Compile("/$")
	if reg2.MatchString(path) {
		path = path + "/"
	}

	for _, file := range files {
		if !file.IsDir() && reg.MatchString(fileName) {
			fileName := path + file.Name()
			list = append(list, &fileName)
		}
	}

	return list
}

func DirRegListDirs(path string, findReg string) []*string {
	list := make([]*string, 0)
	reg, _ := regexp.Compile(findReg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Errorln("open dir failed")
		return nil
	}

	reg2, _ := regexp.Compile("/$")
	if reg2.MatchString(path) {
		path = path + "/"
	}

	for _, file := range files {
		if file.IsDir() && reg.MatchString(fileName) {
			fileName := path + file.Name()
			list = append(list, &fileName)
		}
	}

	return list
}
