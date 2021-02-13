package searcher

import (
	"fmt"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/client"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/slideshare"
	"strings"
	"sync"
)

type searchUseCase struct {
}

func NewSearchUseCase() *searchUseCase {
	return &searchUseCase{}
}

type SearchResponse struct {
	Feedly     []Feed
	SlideShare []Slide
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

type Slide struct {
	ID            int
	Title         string
	Description   string
	URL           string
	ImageURL      string
	EmbedURL      string
	DownloadURL   string
	DownloadCount int
}

const searchCount = 30
const searchPage = 1

func (s searchUseCase) Execute(q string) SearchResponse {
	queries := strings.Split(q, ",")
	if len(queries) == 0 {
		return SearchResponse{}
	}

	fRes := s.searchFeedly(queries)
	sRes := s.searchSlide(queries)

	return SearchResponse{
		Feedly:     fRes,
		SlideShare: sRes,
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

func (s searchUseCase) searchSlide(queries []string) []Slide {
	lock := sync.Mutex{}
	var results []slideshare.SearchSlideResponse
	var wg sync.WaitGroup
	ch := make(chan struct{}, len(queries))

	c := client.NewSlideShareClient()
	c.Init()

	for _, query := range queries {
		fmt.Printf("%s\n", query)
		wg.Add(1)
		go func(query string) {
			defer wg.Done()
			ch <- struct{}{}

			res := c.Search(query, searchCount, searchPage)
			lock.Lock()
			results = append(results, res)
			lock.Unlock()
			<-ch
		}(query)
	}
	wg.Wait()
	close(ch)

	var res []Slide
	for _, v := range results {
		for _, r := range v.Results {
			res = append(res, Slide{
				ID:            r.ID,
				Title:         r.Title,
				Description:   r.Description,
				URL:           r.URL,
				ImageURL:      r.ThumbnailURL,
				EmbedURL:      r.Embed,
				DownloadURL:   r.DownloadURL,
				DownloadCount: r.Download,
			})
		}
	}
	return res
}
