package searcher

import (
	"github.com/gin-gonic/gin"
	searcher "github.com/shinjiezumi/vue-go-samples/src/api/searcher/driver"
	"net/http"
	"strconv"
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

	d := searcher.NewFeedlyDriver()
	d.Init()
	result := d.Search(q, count, page)

	res := make([]searchResponse, 0)
	for _, r := range result.Results {
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

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
