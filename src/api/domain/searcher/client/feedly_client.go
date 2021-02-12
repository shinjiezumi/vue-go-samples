package client

import (
	"encoding/json"
	"fmt"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FeedlyClient struct {
	endpoint string
	apiToken string
}

func NewFeedlyClient() *FeedlyClient {
	return &FeedlyClient{}
}

func (d *FeedlyClient) Init() {
	d.endpoint = "https://cloud.feedly.com/v3"
	d.apiToken = os.Getenv("FEEDLY_ACCESS_TOKEN")
}

func (d *FeedlyClient) Search(keyword string, count, page int) feedly.SearchFeedResponse {
	var result feedly.SearchFeedResponse
	if keyword == "" {
		log.Println("keyword is empty")
		return result
	}
	log.Println("search start")

	url := fmt.Sprintf("%s/?query=%s&locale=ja&count=%d&page=%d", d.endpoint+"/search/feeds", keyword, count, page)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "OAuth "+d.apiToken)
	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		panic(err)
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

	log.Println("search end")
	return result
}
