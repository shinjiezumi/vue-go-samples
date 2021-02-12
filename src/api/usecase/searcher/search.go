package searcher

import (
	"fmt"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/client"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"strings"
	"sync"
)

type searchUseCase struct {
}

func NewSearchUseCase() *searchUseCase {
	return &searchUseCase{}
}

type SearchResponse struct {
	Feedly []Feed
}

type Feed struct {
	ID          string
	Title       string
	Description string
	URL         string
	Subscribers int
	Velocity    float32
	ImageURL    string
	Tags        []string
}

const searchCount = 100
const searchPage = 1

func (s searchUseCase) Execute(q string) SearchResponse {
	queries := strings.Split(q, ",")
	if len(queries) == 0 {
		return SearchResponse{}
	}

	feedlyRes := s.searchFeedly(queries)

	return SearchResponse{
		Feedly: feedlyRes,
	}
}

func (s searchUseCase) searchFeedly(queries []string) []Feed {
	lock := sync.Mutex{}
	var results []feedly.SearchFeedResponse
	var wg sync.WaitGroup
	ch := make(chan struct{}, len(queries))

	d := client.NewFeedlyClient()
	d.Init()

	for _, query := range queries {
		fmt.Printf("%s\n", query)
		wg.Add(1)
		go func(query string) {
			defer wg.Done()
			ch <- struct{}{}

			res := d.Search(query, searchCount, searchPage)
			lock.Lock()
			results = append(results, res)
			lock.Unlock()
			<-ch
		}(query)
	}
	wg.Wait()
	close(ch)

	var res []Feed
	for _, v := range results {
		for _, r := range v.Results {
			res = append(res, Feed{
				ID:          r.FeedID,
				Title:       r.Title,
				Description: r.GetDescription(),
				URL:         r.GetSiteURL(),
				Subscribers: r.Subscribers,
				Velocity:    r.GetVelocity(),
				ImageURL:    r.GetSiteImageURL(),
				Tags:        r.DeliciousTags,
			})
		}
	}
	return res
}
