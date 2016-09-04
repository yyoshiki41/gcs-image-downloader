package downloader

import (
	"log"
	"net/url"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath string
	q        string
)

func Run(args []string) {
	app := kingpin.New("gcs-img-dl", "Image downloader for Google Custom Search API.")

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
}

type GcsAPI struct {
	url.URL
}

func NewGcsAPI() *GcsAPI {
	g := GcsAPI{}
	g.Scheme = "https"
	g.Host = "www.googleapis.com"
	g.Path = "customsearch/v1"
	return &g
}

func (g *GcsAPI) setConfig(c Config) {
	query := g.Query()
	query.Set("cx", c.API.Cx)
	query.Set("key", c.API.Key)
	query.Set("searchType", "image")
	g.RawQuery = query.Encode()
}

func (g *GcsAPI) setQuery(word string) {
	query := g.Query()
	query.Set("q", word)
	g.RawQuery = query.Encode()
}
