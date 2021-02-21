package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
)

const searchFeedCount = 100

// IFeedlyClient FeedlyClientインターフェース
type IFeedlyClient interface {
	Init()
	Search(keyword string) (*feedly.SearchResults, error)
}

// FeedlyClient FeedlyClient
type FeedlyClient struct {
	endpoint string
	apiToken string
}

// NewFeedlyClient APIClientを生成する
func NewFeedlyClient() IFeedlyClient {
	return &FeedlyClient{}
}

// Init APIClientを初期化する
func (c *FeedlyClient) Init() {
	c.endpoint = "https://cloud.feedly.com/v3"
	c.apiToken = os.Getenv("FEEDLY_ACCESS_TOKEN")
}

// Search フィードを検索する
func (c *FeedlyClient) Search(keyword string) (*feedly.SearchResults, error) {
	if keyword == "" {
		return nil, common.NewApplicationError(http.StatusBadRequest, common.InvalidRequest, nil)
	}

	// URL生成
	u, err := c.makeURL(keyword)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	// 検索実行
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	req.Header.Set("Authorization", "OAuth "+c.apiToken)
	hc := new(http.Client)
	res, err := hc.Do(req)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	// レスポンス生成
	ret, err := c.makeResponse(res)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	return ret, nil
}

func (c *FeedlyClient) makeURL(keyword string) (*url.URL, error) {
	// URL生成
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}
	u.Path = path.Join(u.Path, "/search/feeds")

	// クエリ生成
	q := u.Query()
	q.Set("query", keyword)
	q.Set("count", strconv.Itoa(searchFeedCount))
	q.Set("locale", "ja")
	u.RawQuery = q.Encode()

	return u, nil
}

func (c *FeedlyClient) makeResponse(res *http.Response) (*feedly.SearchResults, error) {
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var w feedly.SearchResponseWrapper
	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}

	ret := w.GetSearchResults()
	return &ret, nil
}
