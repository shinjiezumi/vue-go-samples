package feedly

import (
	"fmt"
	"net/url"
)

// SearchResponseWrapper 検索結果ラッパー
type SearchResponseWrapper struct {
	Results SearchResults `json:"results"`
}

// GetSearchResults 検索結果一覧を返す
func (w *SearchResponseWrapper) GetSearchResults() SearchResults {
	return w.Results
}

// SearchResults 検索結果一覧
type SearchResults []SearchResult

// SearchResult 検索結果
type SearchResult struct {
	FeedID        string   `json:"feedId"`        // フィードID
	Subscribers   int      `json:"subscribers"`   // 登録者数
	Score         float32  `json:"score"`         // スコア
	Title         string   `json:"title"`         // サイト名
	Description   *string  `json:"description"`   // サイト概要
	Website       *string  `json:"website"`       // ウェブサイトURL
	LastUpdated   *int     `json:"lastUpdated"`   // 最終更新タイムスタンプ
	Velocity      *float32 `json:"velocity"`      // 週次投稿数
	Language      *string  `json:"language"`      // 言語
	Featured      *bool    `json:"featured"`      // オススメ
	IconURL       *string  `json:"iconUrl"`       // 小アイコン画像URL
	VisualURL     *string  `json:"visualUrl"`     // 大アイコン画像URL
	CoverURL      *string  `json:"coverUrl"`      // 背景画像URL
	Logo          *string  `json:"logo"`          // ロゴ
	ContentType   *string  `json:"contentType"`   // コンテンツタイプ
	CoverColor    *string  `json:"coverColor"`    // 背景色(16進数)
	DeliciousTags []string `json:"deliciousTags"` // タグ
}

// GetSiteURL サイトURLを返す
func (r *SearchResult) GetSiteURL() string {
	if r.Website != nil {
		return *r.Website
	}

	return ""
}

// GetDescription サイト概要を返す
func (r *SearchResult) GetDescription() string {
	if r.Description != nil {
		return *r.Description
	}

	return ""
}

// GetSiteImageURL サイト画像URLを返す
func (r *SearchResult) GetSiteImageURL() string {
	if r.VisualURL != nil {
		return *r.VisualURL
	}

	return fmt.Sprintf("https://placehold.jp/%dx%d.png?text=%s", 150, 150, url.QueryEscape("NO IMAGE"))
}
