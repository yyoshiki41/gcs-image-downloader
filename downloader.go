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
	"sync/atomic"

	"github.com/yyoshiki41/gcs-image-downloader/internal/entity"
	"github.com/yyoshiki41/gcs-image-downloader/internal/file"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath    string
	outputsPath string
	q           string
	num         int
)

func Run(args []string) {
	app := kingpin.New("downloader", "Image downloader for Google Custom Search API.")

	app.Flag("conf", "Config file path").Default("conf").Short('c').StringVar(&confPath)
	app.Flag("outputs", "Outputs directory").Default("outputs").Short('o').StringVar(&outputsPath)
	app.Flag("query", "Query").Required().Short('q').StringVar(&q)
	app.Flag("number", "Number of files").Default("10").Short('n').IntVar(&num)

	kingpin.MustParse(app.Parse(args))

	if _, err := os.Stat(outputsPath); err != nil {
		log.Fatalf("%v\noutputs: %v", err, outputsPath)
	}

	var conf Config
	err := loadConf(confPath, &conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start!")
	fmt.Printf("Number: %v\n", num)

	var results []entity.Link
	var wg sync.WaitGroup

	n := num / 10
	for i := 0; i < n; i++ {
		index := 10*i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()

			r := run(conf, index)
			results = append(results, r.Items...)
		}()
	}
	wg.Wait()

	var errCount int64
	for _, v := range results {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			err := download(link)
			if err != nil {
				atomic.AddInt64(&errCount, 1)
				log.Println(err)
			}
		}(v.Link)
	}
	wg.Wait()

	total := len(results)
	log.Println("Download has completed!")
	fmt.Printf("Total: %v, Success: %v, Failure: %v\n", total, int64(total)-errCount, errCount)
}

func run(conf Config, index int) *entity.GcsResponse {
	gcs := NewGcsAPI()
	gcs.setConfig(conf)
	gcs.setStart(index)
	gcs.setQuery(q)
	b := gcs.Get()

	resp := entity.NewGcsResponse()
	json.Unmarshal(b, &resp)
	if resp == nil {
		log.Println("CustomSearchAPI Response is empty.")
	}
	return resp
}

func download(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(path.Join(outputsPath, file.Name(link)))
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return err
}
