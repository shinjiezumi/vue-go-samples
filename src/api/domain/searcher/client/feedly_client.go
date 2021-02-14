package client

import (
	"encoding/json"
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

type FeedlyClient struct {
	endpoint string
	apiToken string
}

func NewFeedlyClient() *FeedlyClient {
	return &FeedlyClient{}
}

func (c *FeedlyClient) Init() {
	c.endpoint = "https://cloud.feedly.com/v3"
	c.apiToken = os.Getenv("FEEDLY_ACCESS_TOKEN")
}

func (c *FeedlyClient) Search(keyword string, count, page int) (*feedly.SearchFeedResponse, error) {
	var result feedly.SearchFeedResponse
	if keyword == "" {
		return nil, common.NewApplicationError(http.StatusBadRequest, common.InvalidRequest)
	}

	// URL生成
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
	u.Path = path.Join(u.Path, "/search/feeds")

	// クエリ生成
	q := u.Query()
	q.Set("query", url.QueryEscape(keyword))
	q.Set("page", strconv.Itoa(page))
	q.Set("page", strconv.Itoa(count))
	q.Set("locale", "ja")
	u.RawQuery = q.Encode()

	// 検索実行
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}

	req.Header.Set("Authorization", "OAuth "+c.apiToken)
	hc := new(http.Client)
	res, err := hc.Do(req)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}

	return &result, nil
}
