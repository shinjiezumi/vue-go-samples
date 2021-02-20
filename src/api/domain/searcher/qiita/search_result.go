package qiita

import "time"

// SearchResults 検索結果一覧
type SearchResults []SearchResult

// SearchResult 検索結果
type SearchResult struct {
	RenderedBody   string    `json:"rendered_body"`    // 本文(HTML)
	Body           string    `json:"body"`             // 本文(Markdown)
	Coediting      *bool     `json:"coediting"`        // 共同更新かどうか(QiitaTeamのみ有効)
	CommentsCount  int       `json:"comments_count"`   // コメント数
	CreatedAt      time.Time `json:"created_at"`       // 作成日時
	Group          string    `json:"group"`            // QiitaTeamのグループ
	ID             string    `json:"id"`               // ID
	LikesCount     int       `json:"likes_count"`      // LGTM数(Qiitaのみ有効)
	Private        bool      `json:"private"`          // 限定共有状態かどうか(Qiita Teamでは無効)
	ReactionsCount int       `json:"reactions_count"`  // 絵文字リアクション数(Qiita Teamのみ有効)
	Tags           Tags      `json:"tags"`             // タグ
	Title          string    `json:"title"`            // タイトル
	UpdatedAt      time.Time `json:"updated_at"`       // 更新日時
	URL            string    `json:"url"`              // URL
	User           User      `json:"-"`                // ユーザー情報 ※割愛
	PageViewsCount int       `json:"page_views_count"` // PV数 ※死んでるよう
}

// Tag タグ
type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// Tags タグ一覧
type Tags []Tag

// GetTags タグ一覧を返す
func (t *Tags) GetTags() []string {
	ret := make([]string, 0, len(*t))

	for _, v := range *t {
		ret = append(ret, v.Name)
	}

	return ret
}

// User ユーザー情報(実装割愛)
type User struct{}
