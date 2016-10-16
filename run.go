package downloader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/yyoshiki41/gcs-image-downloader/internal/entity"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath    string
	outputsPath string
	q           string
	num         int
	safe        string
	imgType     string
)

func Run(args []string) {
	app := kingpin.New("downloader", "Image downloader for Google Custom Search API.")

	app.Flag("conf", "Config file path").Default("conf").Short('c').StringVar(&confPath)
	app.Flag("outputs", "Outputs directory").Default("outputs").Short('o').StringVar(&outputsPath)
	app.Flag("query", "Query").Required().Short('q').StringVar(&q)
	app.Flag("number", "Number of files").Default("10").Short('n').IntVar(&num)
	app.Flag("safe", "Safety level: high, medium, off").PlaceHolder("SAFETY-LEVEL").Short('s').StringVar(&safe)
	app.Flag("type", "Images of a type: clipart, face, lineart, news, photo").PlaceHolder("IMG-TYPE").Short('t').StringVar(&imgType)

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

	n := (num + 9) / 10
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

	errCount := bulkDownload(results)

	log.Println("Download has completed!")
	fmt.Printf("Total: %v, Success: %v, Failure: %v\n", num, int64(num)-errCount, errCount)
}

func run(conf Config, index int) *entity.GcsResponse {
	gcs := NewGcsAPI()
	gcs.setConfig(conf)
	gcs.setStart(index)
	gcs.setQuery(q)
	gcs.setSafe(safe)
	gcs.setImgType(imgType)
	b := gcs.Get()

	resp := entity.NewGcsResponse()
	json.Unmarshal(b, &resp)
	if resp == nil {
		log.Println("CustomSearchAPI Response is empty.")
	}
	return resp
}
