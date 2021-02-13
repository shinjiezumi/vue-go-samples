package client

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
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

func (c *SlideShareClient) Search(keyword string, count, page int) slideshare.SearchSlideResponse {
	var ret slideshare.SearchSlideResponse
	if keyword == "" {
		log.Println("keyword is empty")
		return ret
	}
	log.Println("search start")

	u, err := url.Parse(c.endpoint)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, "search_slideshows")

	q := u.Query()
	q.Set("q", url.QueryEscape(keyword))
	q.Set("page", strconv.Itoa(page))
	q.Set("items_per_page", strconv.Itoa(count))
	q = c.addCommonQuery(q)

	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	hc := new(http.Client)
	res, err := hc.Do(req)
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

	err = xml.Unmarshal(body, &ret)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("search end")
	return ret
}

func (c *SlideShareClient) addCommonQuery(q url.Values) url.Values {
	// 認証情報設定
	q.Set("api_key", c.apikey)
	q.Set("ts", strconv.FormatInt(c.ts, 10))
	q.Set("hash", c.hash)
	q.Set("lang", "ja")

	return q
}
