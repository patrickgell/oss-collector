package io

import (
	"fmt"
	"github.com/pat-git023/oss-collector/v2/internal/model"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadSourcesAndCreateBigZIP(project *model.Project) {
	components := project.Components
	for _, component := range components {
		filename := downloadUrl(&component)
		if len(filename) > 0 {
			Unzip(filename, "sources/")
		}
	}
	CreateZip("sources/", project.Name+".zip")
}

func downloadUrl(component *model.Components) string {
	tokens := strings.Split(component.SourcesURL, "/")
	fileName := tokens[len(tokens)-1]
	log.Printf("Downloading %s to %s\n", component.SourcesURL, fileName)

	exist := exists(fileName)
	if exist == true {
		log.Printf("file %s already exists - skipping\n", fileName)
		return fileName
	}
	output, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error while creating %s: %v", fileName, err)
	}
	defer output.Close()

	response, err := http.Get(component.SourcesURL)
	if err != nil {
		log.Fatalf("Error while downloading %s - %v", component.SourcesURL, err)
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Fatalf("Error while downloading %s - %v", component.SourcesURL, err)
	}

	fmt.Println(n, "bytes downloaded.")
	return fileName

}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
