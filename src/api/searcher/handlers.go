package searcher

import (
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/usecase/searcher"
	"net/http"
	"strconv"
)

// Search は検索して結果を返します
func Search(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": []searcher.SearchResponse{},
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

	res := searcher.NewSearchUseCase().Execute(q, count, page)

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
