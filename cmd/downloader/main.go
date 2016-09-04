package main

import (
	"os"

	gcs "github.com/yyoshiki41/gcs-image-downloader"
)

func main() {
	gcs.Run(os.Args[1:])
}
