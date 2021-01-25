package searcher

import (
	"encoding/json"
	"fmt"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FeedlyDriver struct {
	endpoint string
	apiToken string
}

func NewFeedlyDriver() *FeedlyDriver {
	return &FeedlyDriver{}
}

func (d *FeedlyDriver) Init() {
	d.endpoint = "https://cloud.feedly.com/v3"
	d.apiToken = os.Getenv("FEEDLY_ACCESS_TOKEN")
}

func (d *FeedlyDriver) Search(keyword string, count, page int) feedly.SearchResponse {
	var result feedly.SearchResponse
	if keyword == "" {
		fmt.Println("keyword is empty")
		return result
	}
	fmt.Println("search start")

	url := fmt.Sprintf("%s/?query=%s&locale=ja&count=%d&page=%d", d.endpoint+"/search/feeds", keyword, count, page)
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		log.Fatal(err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("search end")
	return result
}
