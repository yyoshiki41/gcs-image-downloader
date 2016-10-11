package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
	query.Set("fields", "items/link")
	g.RawQuery = query.Encode()
}

func (g *GcsAPI) setQuery(word string) {
	query := g.Query()
	query.Set("q", word)
	g.RawQuery = query.Encode()
}

func (g *GcsAPI) Get() []byte {
	resp, err := http.Get(g.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
