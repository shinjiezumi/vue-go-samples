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
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/qiita"
)

const searchQiitaCount = 80
const searchQiitaPage = 1

// IQiitaClient QiitaAPIClientインターフェース
type IQiitaClient interface {
	Init()
	Search(keyword string) (*qiita.SearchResults, error)
}

// QiitaClient QiitaAPIClient
type QiitaClient struct {
	endpoint string
	apiToken string
}

// NewQiitaClient APIClientを生成する
func NewQiitaClient() IQiitaClient {
	return &QiitaClient{}
}

// Init APIClientを初期化する
func (c *QiitaClient) Init() {
	c.endpoint = "https://qiita.com/api/v2"
	c.apiToken = os.Getenv("QIITA_API_TOKEN")
}

// Search Qiita記事を検索する
func (c *QiitaClient) Search(keyword string) (*qiita.SearchResults, error) {
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
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	hc := new(http.Client)
	res, err := hc.Do(req)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	// レスポンス生成
	ret, err := c.makeResponse(res)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	return ret, nil
}

func (c *QiitaClient) makeURL(keyword string) (*url.URL, error) {
	// URL生成
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "items")

	// クエリ生成
	q := u.Query()
	q.Set("query", keyword)
	q.Set("page", strconv.Itoa(searchQiitaPage))
	q.Set("per_page", strconv.Itoa(searchQiitaCount))
	u.RawQuery = q.Encode()

	return u, nil
}

func (c *QiitaClient) makeResponse(res *http.Response) (*qiita.SearchResults, error) {
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ret qiita.SearchResults
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
