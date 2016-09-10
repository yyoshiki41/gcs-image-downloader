package downloader

import (
	"fmt"
	"log"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath string
	q        string
)

func Run(args []string) {
	app := kingpin.New("downloader", "Image downloader for Google Custom Search API.")

	app.Flag("conf", "config file path").Default("conf").Short('c').StringVar(&confPath)
	app.Flag("query", "query").Required().Short('q').StringVar(&q)

	kingpin.MustParse(app.Parse(args))

	var conf Config
	err := loadConf(confPath, &conf)
	if err != nil {
		log.Fatal(err)
	}

	gcs := NewGcsAPI()
	gcs.setConfig(conf)
	gcs.setQuery(q)
	resp := gcs.Get()
	fmt.Println(string(resp))
}
