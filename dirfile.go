package gotool

import (
	"io/ioutil"
	"regexp"
)

func DirRegListFiles(path string, findReg string) ([]*string, error) {
	list := make([]*string, 0)
	reg, _ := regexp.Compile(findReg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	reg2, _ := regexp.Compile("/$")
	if !reg2.MatchString(path) {
		path = path + "/"
	}

	for _, file := range files {
		if !file.IsDir() && reg.MatchString(file.Name()) {
			fileName := path + file.Name()
			list = append(list, &fileName)
		}
	}

	return list, nil
}

func DirRegListDirs(path string, findReg string) ([]*string, error) {
	list := make([]*string, 0)
	reg, _ := regexp.Compile(findReg)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	reg2, _ := regexp.Compile("/$")
	if !reg2.MatchString(path) {
		path = path + "/"
	}

	for _, file := range files {
		if file.IsDir() && reg.MatchString(file.Name()) {
			fileName := path + file.Name()
			list = append(list, &fileName)
		}
	}

	return list, nil
}
