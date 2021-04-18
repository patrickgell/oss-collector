package main

import "github.com/pat-git023/oss-collector/v2/internal/io"

func main() {

	// read components
	project := io.ReadJsonFile("oss-components.json")
	// download sources and create ZIP
	io.DownloadSourcesAndCreateBigZIP(project)
}
