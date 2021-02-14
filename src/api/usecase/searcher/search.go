package searcher

import (
	"context"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/client"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/feedly"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/qiita"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/searcher/slideshare"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
)

type searchUseCase struct{}

func NewSearchUseCase() *searchUseCase {
	return &searchUseCase{}
}

type SearchResponse struct {
	Feedly     []Feed
	SlideShare []Slide
	Qiita      []QiitaItem
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

type QiitaItem struct {
	ID        string
	Title     string
	LikeCount int
	Tags      []string
	URL       string
	PVCount   int
}

const searchCount = 30
const searchPage = 1

func (s searchUseCase) Execute(q string) SearchResponse {
	queries := strings.Split(q, ",")
	if len(queries) == 0 {
		return SearchResponse{}
	}

	// TODO エラー返す
	fRes, _ := s.searchFeedly(queries)
	sRes, _ := s.searchSlide(queries)
	qRes, _ := s.searchQiitaItem(queries)

	return SearchResponse{
		Feedly:     fRes,
		SlideShare: sRes,
		Qiita:      qRes,
	}
}

func (s searchUseCase) searchFeedly(queries []string) ([]Feed, error) {
	lock := sync.Mutex{}
	var results []feedly.SearchFeedResponse

	eg, ctx := errgroup.WithContext(context.Background())
	d := client.NewFeedlyClient()
	d.Init()
	for _, query := range queries {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := d.Search(query, searchCount, searchPage)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, *res)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

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
	return res, nil
}

func (s searchUseCase) searchSlide(queries []string) ([]Slide, error) {
	lock := sync.Mutex{}
	var results []slideshare.SearchSlideResponse

	eg, ctx := errgroup.WithContext(context.Background())
	c := client.NewSlideShareClient()
	c.Init()
	for _, query := range queries {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := c.Search(query, searchCount, searchPage)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, *res)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

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
	return res, nil
}

func (s searchUseCase) searchQiitaItem(queries []string) ([]QiitaItem, error) {
	lock := sync.Mutex{}
	var results []qiita.SearchItemResponse

	eg, ctx := errgroup.WithContext(context.Background())
	c := client.NewQiitaClient()
	c.Init()
	for _, query := range queries {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := c.Search(query, searchCount, searchPage)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, res)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	var res []QiitaItem
	for _, v := range results {
		for _, r := range v {
			res = append(res, QiitaItem{
				ID:        r.ID,
				Title:     r.Title,
				LikeCount: r.LikesCount,
				Tags:      r.Tags.GetTags(),
				URL:       r.URL,
				PVCount:   r.PageViewsCount,
			})
		}
	}
	return res, nil
}
