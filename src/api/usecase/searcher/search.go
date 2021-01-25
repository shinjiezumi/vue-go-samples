package searcher

import (
	"fmt"
	searcher "github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/driver"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"strings"
	"sync"
)

type SearchUseCase struct {
}

func NewSearchUseCase() *SearchUseCase {
	return &SearchUseCase{}
}

type SearchResponse struct {
	FeedID      string
	Title       string
	Description string
	Thumbnail   string
	URL         string
	Tags        []string
	Subscribers int
}

func (s SearchUseCase) Execute(q string, count, page int) []SearchResponse {
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

	var res []SearchResponse
	for _, v := range result {
		for _, r := range v.Results {
			res = append(res, SearchResponse{
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
	return res
}
