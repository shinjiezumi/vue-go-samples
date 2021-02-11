package searcher

import (
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/usecase/searcher"
	"net/http"
)

// Search は検索して結果を返します
func Search(c *gin.Context) {
	q := c.DefaultQuery("q", "")

	res := searcher.NewSearchUseCase().Execute(q)

	c.JSON(http.StatusOK, res)
}
