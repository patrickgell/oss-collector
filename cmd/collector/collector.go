package main

import (
	"flag"
	"github.com/pat-git023/oss-collector/v2/internal/io"
)

var (
	configFile = flag.String("config", "oss-components.json", "Path to the config file")
)

func main() {
	flag.Parse()

	// read components
	project := io.ReadJsonFile(configFile)
	// download sources and create ZIP
	io.DownloadSourcesAndCreateBigZIP(project)
}
