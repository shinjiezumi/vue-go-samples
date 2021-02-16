package client

import (
	"encoding/json"
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/qiita"
	"io/ioutil"
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

func (c *QiitaClient) Search(keyword string, count, page int) (*qiita.SearchResponse, error) {
	var ret qiita.SearchResponse
	if keyword == "" {
		return nil, common.NewApplicationError(http.StatusBadRequest, common.InvalidRequest)
	}

	// URL生成
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
	u.Path = path.Join(u.Path, "items")

	// クエリ生成
	q := u.Query()
	q.Set("query", url.QueryEscape(keyword))
	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(count))
	u.RawQuery = q.Encode()

	// 検索実行
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
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

	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}

	return &ret, nil
}
