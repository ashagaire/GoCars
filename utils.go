package main

import (
	"io/ioutil"
)

func getTemplateFiles() []string {
	templateFiles := []string{}
	files, filesErr := ioutil.ReadDir("./templates")
	if filesErr != nil {
		return templateFiles
	}

	for _, file := range files {
		name := "templates/" + file.Name()
		templateFiles = append(templateFiles, name)
	}
	return templateFiles
}
