package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

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
		var wg sync.WaitGroup
		for _, v := range resp.Items {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				download(link)
			}(v.Link)
		}
		wg.Wait()
	}
}

func download(link string) error {
	fmt.Println(link)

	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, fileName := path.Split(link)
	file, err := os.Create(path.Join(outputsPath, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}
