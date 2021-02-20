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
	"time"
)

type searchUseCase struct{}

func NewSearchUseCase() *searchUseCase {
	return &searchUseCase{}
}

type SearchResponse struct {
	Feedly     FeedlyResponse
	SlideShare SlideShareResponse
	Qiita      QiitaResponse
}

type SearchError struct {
	Message string
}

type FeedlyResponse struct {
	List  []Feed
	Error SearchError
}

type SlideShareResponse struct {
	List  []Slide
	Error SearchError
}

type QiitaResponse struct {
	List  []Qiita
	Error SearchError
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

type Qiita struct {
	ID        string
	Title     string
	LikeCount int
	Tags      []string
	URL       string
}

const searchExpirySecond = 10

func (s searchUseCase) Execute(q string) SearchResponse {
	queries := strings.Split(q, ",")
	if len(queries) == 0 {
		return SearchResponse{}
	}

	var fRes FeedlyResponse
	var sRes SlideShareResponse
	var qRes QiitaResponse
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fRes = s.searchFeedly(queries)
	}()
	go func() {
		defer wg.Done()
		sRes = s.searchSlide(queries)
	}()
	go func() {
		defer wg.Done()
		qRes = s.searchQiita(queries)
	}()
	wg.Wait()

	return SearchResponse{
		Feedly:     fRes,
		SlideShare: sRes,
		Qiita:      qRes,
	}
}

func (s searchUseCase) searchFeedly(queries []string) FeedlyResponse {
	// コンテキスト設定
	pCtx, cancel := context.WithTimeout(context.Background(), searchExpirySecond*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(pCtx)

	// client生成
	d := client.NewFeedlyClient()
	d.Init()

	// 検索
	lock := sync.Mutex{}
	var results feedly.SearchResults
	for _, query := range queries {
		q := query
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := d.Search(q)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, *res...)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return FeedlyResponse{
			Error: SearchError{
				err.Error(),
			},
		}
	}

	// ソート
	sorted := feedly.Sort(results)

	var res []Feed
	for _, r := range sorted {
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

	return FeedlyResponse{
		List: res,
	}
}

func (s searchUseCase) searchSlide(queries []string) SlideShareResponse {
	// コンテキスト設定
	pCtx, cancel := context.WithTimeout(context.Background(), searchExpirySecond*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(pCtx)

	// client生成
	c := client.NewSlideShareClient()
	c.Init()

	// 検索
	lock := sync.Mutex{}
	var results slideshare.SearchResults
	for _, query := range queries {
		q := query
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := c.Search(q)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, *res...)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return SlideShareResponse{
			Error: SearchError{
				err.Error(),
			},
		}
	}

	sorted := slideshare.Sort(results)

	var res []Slide
	for _, r := range sorted {
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

	return SlideShareResponse{
		List: res,
	}
}

func (s searchUseCase) searchQiita(queries []string) QiitaResponse {
	// コンテキスト設定
	pCtx, cancel := context.WithTimeout(context.Background(), searchExpirySecond*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(pCtx)

	//  client生成
	c := client.NewQiitaClient()
	c.Init()

	// 検索
	lock := sync.Mutex{}
	var results qiita.SearchResults
	for _, query := range queries {
		q := query
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := c.Search(q)
				if err != nil {
					return err
				}
				lock.Lock()
				results = append(results, *res...)
				lock.Unlock()
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return QiitaResponse{
			Error: SearchError{
				err.Error(),
			},
		}
	}

	// ソート
	sorted := qiita.Sort(results)

	var res []Qiita
	for _, r := range sorted {
		res = append(res, Qiita{
			ID:        r.ID,
			Title:     r.Title,
			LikeCount: r.LikesCount,
			Tags:      r.Tags.GetTags(),
			URL:       r.URL,
		})
	}

	return QiitaResponse{
		List: res,
	}
}
