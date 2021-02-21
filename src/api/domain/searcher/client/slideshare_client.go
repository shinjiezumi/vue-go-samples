package client

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/slideshare"
)

const searchSlideCount = 50
const searchSlidePage = 1

// ISlideShareClient ISlideShareClientインターフェース
type ISlideShareClient interface {
	Init()
	Search(keyword string) (*slideshare.SearchResults, error)
}

// SlideShareClient SlideShareClient
type SlideShareClient struct {
	endpoint string
	apikey   string
	ts       int64
	hash     string
}

// NewSlideShareClient APIClientを生成する
func NewSlideShareClient() ISlideShareClient {
	return &SlideShareClient{}
}

// Init APIClientを初期化する
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

// Search スライドを検索する
func (c *SlideShareClient) Search(keyword string) (*slideshare.SearchResults, error) {
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

	ret, err := c.makeResponse(res)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	return ret, nil
}

func (c *SlideShareClient) addCommonQuery(q url.Values) url.Values {
	q.Set("api_key", c.apikey)
	q.Set("ts", strconv.FormatInt(c.ts, 10))
	q.Set("hash", c.hash)
	q.Set("lang", "ja")

	return q
}

func (c *SlideShareClient) makeURL(keyword string) (*url.URL, error) {
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "search_slideshows")

	q := u.Query()
	q.Set("q", keyword)
	q.Set("page", strconv.Itoa(searchSlidePage))
	q.Set("items_per_page", strconv.Itoa(searchSlideCount))
	q = c.addCommonQuery(q)
	u.RawQuery = q.Encode()

	return u, nil
}

func (c *SlideShareClient) makeResponse(res *http.Response) (*slideshare.SearchResults, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	var w slideshare.SearchResponseWrapper
	err = xml.Unmarshal(body, &w)
	if err != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, common.GeneralError, err)
	}

	ret := w.GetSearchResults()

	return &ret, nil
}
