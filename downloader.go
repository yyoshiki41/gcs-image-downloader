package downloader

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yyoshiki41/gcs-image-downloader/internal/entity"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath    string
	outputsPath string
	q           string
)

func Run(args []string) {
	app := kingpin.New("downloader", "Image downloader for Google Custom Search API.")

	app.Flag("conf", "config file path").Default("conf").Short('c').StringVar(&confPath)
	app.Flag("outputs", "Directory for downloaded files").Default("outputs").Short('o').StringVar(&outputsPath)
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
	b := gcs.Get()

	resp := new(entity.GcsResponse)
	json.Unmarshal(b, &resp)

	if resp != nil {
		for _, v := range resp.Items {
			fmt.Println(v.Link)
		}
	}
}
