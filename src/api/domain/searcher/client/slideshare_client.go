package client

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/slideshare"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

type SlideShareClient struct {
	endpoint string
	apikey   string
	ts       int64
	hash     string
}

func NewSlideShareClient() *SlideShareClient {
	return &SlideShareClient{}
}

func (c *SlideShareClient) Init() {
	c.endpoint = "https://www.slideshare.net/api/2"
	// @see https://www.slideshare.net/developers/documentation
	c.apikey = os.Getenv("SLIDE_SHARE_API_KEY")
	c.ts = time.Now().Unix()
	secret := os.Getenv("SLIDE_SHARE_API_SECRET")
	s := sha1.New()
	_, err := io.WriteString(s, secret+strconv.FormatInt(c.ts, 10))
	if err != nil {
		log.Fatal(err)
	}
	c.hash = hex.EncodeToString(s.Sum(nil))
}

func (c *SlideShareClient) Search(keyword string, count, page int) (*slideshare.SearchSlideResponse, error) {
	var ret slideshare.SearchSlideResponse
	if keyword == "" {
		return nil, common.NewApplicationError(http.StatusBadRequest, common.InvalidRequest)
	}

	// URL生成
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
	u.Path = path.Join(u.Path, "search_slideshows")

	// クエリ生成
	q := u.Query()
	q.Set("q", url.QueryEscape(keyword))
	q.Set("page", strconv.Itoa(page))
	q.Set("items_per_page", strconv.Itoa(count))
	q = c.addCommonQuery(q)
	u.RawQuery = q.Encode()

	// 検索実行
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}
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

	err = xml.Unmarshal(body, &ret)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError)
	}

	return &ret, nil
}

func (c *SlideShareClient) addCommonQuery(q url.Values) url.Values {
	// 認証情報設定
	q.Set("api_key", c.apikey)
	q.Set("ts", strconv.FormatInt(c.ts, 10))
	q.Set("hash", c.hash)
	q.Set("lang", "ja")

	return q
}
