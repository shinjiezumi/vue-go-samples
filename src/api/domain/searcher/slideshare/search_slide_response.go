package slideshare

type SearchSlideResponse struct {
	Results []SearchSlideResult `xml:"Slideshow"`
}

type SearchSlideResult struct {
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
	Download            int           `xml:"Download"`            // ダウンロード数
	DownloadURL         string        `xml:"DownloadURL"`         // ダウンロードURL
	SecretKey           string        `xml:"SecretKey"`           // シークレット
	SlideshowEmbedURL   string        `xml:"SlideshowEmbedURL"`   // スライドショー埋め込みRUL
	SlideshowType       int           `xml:"SlideshowType"`       // スライドショー種別
	InContest           int           `xml:"InContest"`           //
}

type ThumbnailSize []string
