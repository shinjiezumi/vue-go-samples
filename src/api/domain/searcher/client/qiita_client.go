package client

import (
	"encoding/json"
	"fmt"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/qiita"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

type QiitaClient struct {
	endpoint string
	apiToken string
}

func NewQiitaClient() *QiitaClient {
	return &QiitaClient{}
}

func (c *QiitaClient) Init() {
	c.endpoint = "https://qiita.com/api/v2"
	c.apiToken = os.Getenv("QIITA_API_TOKEN")
}

func (c *QiitaClient) Search(keyword string, count, page int) qiita.SearchItemResponse {
	var result qiita.SearchItemResponse
	if keyword == "" {
		log.Println("keyword is empty")
		return result
	}
	log.Println("search start")

	u, err := url.Parse(c.endpoint)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, "items")

	q := u.Query()
	q.Set("query", url.QueryEscape(keyword))
	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(count))

	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	hc := new(http.Client)
	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	log.Println("search end")
	return result
}
