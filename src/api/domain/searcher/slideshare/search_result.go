package slideshare

import "time"

// SearchResponseWrapper 検索結果ラッパー
type SearchResponseWrapper struct {
	Results SearchResults `xml:"Slideshow"`
}

// GetSearchResults 検索結果一覧を返す
func (w *SearchResponseWrapper) GetSearchResults() SearchResults {
	return w.Results
}

// SearchResults 検索結果一覧
type SearchResults []SearchResult

// SearchResult 検索結果
type SearchResult struct {
	ID                  int           `xml:"ID"`                  // スライドID
	Title               string        `xml:"Title"`               // タイトル
	Description         string        `xml:"Description"`         // 概要
	Username            string        `xml:"Username"`            // ユーザー名
	URL                 string        `xml:"URL"`                 // URL
	ThumbnailURL        string        `xml:"ThumbnailURL"`        // サムネイルURL
	ThumbnailSize       ThumbnailSize `xml:"ThumbnailSize"`       // サムネイルサイズ
	ThumbnailSmallURL   string        `xml:"ThumbnailSmallURL"`   // サムネイルURL(小)
	ThumbnailXLargeURL  string        `xml:"ThumbnailXLargeURL"`  // サムネイルURL(大)
	ThumbnailXXLargeURL string        `xml:"ThumbnailXXLargeURL"` // サムネイルURL(特大)
	Embed               string        `xml:"Embed"`               // 埋め込みURL
	Created             string        `xml:"Created"`             // 作成日時
	Updated             string        `xml:"Updated"`             // 更新日時
	Language            string        `xml:"Language"`            // 言語
	Format              string        `xml:"Format"`              // フォーマット
	Download            bool          `xml:"Download"`            // ダウンロード可能フラグ
	DownloadURL         string        `xml:"DownloadURL"`         // ダウンロードURL
	SecretKey           string        `xml:"SecretKey"`           // シークレット
	SlideshowEmbedURL   string        `xml:"SlideshowEmbedURL"`   // スライドショー埋め込みRUL
	SlideshowType       int           `xml:"SlideshowType"`       // スライドショー種別
	InContest           int           `xml:"InContest"`           //
}

func (r *SearchResult) GetUpdated() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05 UTC", r.Updated)

	return t
}

// ThumbnailSize サムネイルサイズ
type ThumbnailSize []string
