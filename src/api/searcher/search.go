package searcher

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type searchResponse struct {
	Q string `json:"q"`
}

// Search は検索結果を取得します
func Search(c *gin.Context) {
	q := c.DefaultQuery("q", "")

	res := searchResponse{
		Q: q,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
