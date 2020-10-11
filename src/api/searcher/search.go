package searcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	searcher "github.com/shinjiezumi/vue-go-samples/src/api/searcher/driver"
	"github.com/shinjiezumi/vue-go-samples/src/api/searcher/feedly"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type searchResponse struct {
	FeedID      string
	Title       string
	Description string
	Thumbnail   string
	URL         string
	Tags        []string
	Subscribers int
}

// Search は検索して結果を返します
func Search(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": []searchResponse{},
		})
		return
	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "20"))
	if err != nil {
		panic(err)
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		panic(err)
	}

	queries := strings.Split(q, ",")
	lock := sync.Mutex{}
	var result []feedly.SearchResponse
	var wg sync.WaitGroup
	ch := make(chan struct{}, 4)
	d := searcher.NewFeedlyDriver()
	d.Init()
	for _, query := range queries {
		fmt.Printf("%s\n", query)
		wg.Add(1)
		go func(query string) {
			defer wg.Done()
			ch <- struct{}{}

			res := d.Search(query, count, page)
			lock.Lock()
			result = append(result, res)
			lock.Unlock()
			<-ch
		}(query)
	}
	wg.Wait()

	res := make([]searchResponse, 0)
	for _, v := range result {
		for _, r := range v.Results {
			res = append(res, searchResponse{
				FeedID:      r.FeedID,
				Title:       r.Title,
				Description: r.Description,
				Thumbnail:   r.VisualURL,
				URL:         r.Website,
				Tags:        r.DeliciousTags,
				Subscribers: r.Subscribers,
			})
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
