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
			err := Unzip(filename, "sources/")
			if err != nil {
				log.Fatalf("could not unzip file %s: %v", filename, err)
			}
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
		return ""
	}
	output, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error while creating %s: %v", fileName, err)
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			log.Fatalf("error during close: %v", err)
		}
	}(output)

	response, err := http.Get(component.SourcesURL)
	if err != nil {
		log.Fatalf("Error while downloading %s - %v", component.SourcesURL, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error during close: %v", err)
		}
	}(response.Body)

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
