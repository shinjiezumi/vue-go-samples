package qiita

import (
	"time"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
)

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
	User           User      `json:"user"`             // ユーザー情報
	PageViewsCount int       `json:"page_views_count"` // PV数 ※死んでるよう
}

// GetTags タグ一覧を返す
func (r *SearchResult) GetTags() []string {
	ret := make([]string, 0, len(r.Tags))

	for _, v := range r.Tags {
		ret = append(ret, v.Name)
	}

	return ret
}

// GetCreatedDate 作成日を返す
func (r *SearchResult) GetCreatedDate() string {
	return r.CreatedAt.Format(common.DateFormat4Show)
}

// GetProfileImageURL プロフィール画像URLを返す
func (r *SearchResult) GetProfileImageURL() string {
	return r.User.ProfileImageURL
}

// Tag タグ
type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// Tags タグ一覧
type Tags []Tag

// User ユーザー情報
type User struct {
	Description       string `json:"description"`       // 自己紹介
	FacebookID        string `json:"facebook_id"`       // FacebookID
	FolloweesCount    int    `json:"followees_count"`   // フォロー数
	FollowersCount    int    `json:"followers_count"`   // フォロワー数,
	GithubLoginName   string `json:"github_login_name"` // GitHubログイン名
	ID                string `json:"id"`                // ユーザーID
	ItemsCount        int    `json:"items_count"`       // 投稿数,
	LinkedinID        string `json:"Linkedin_id"`       //LinkedinID,
	Location          string `json:"location"`          // ロケーション
	Name              string `json:"name"`              // 名前
	Organization      string `json:"organization"`      // 所属
	PermanentID       int    `json:"permanent_id"`
	ProfileImageURL   string `json:"profile_image_url"`   // プロフィール画像URL
	TeamOnly          bool   `json:"team_only"`           //
	TwitterScreenName string `json:"twitter_screen_name"` // ツイッターアカウント名
	WebsiteURL        string `json:"website_url"`         // ウェブサイトURL
}
