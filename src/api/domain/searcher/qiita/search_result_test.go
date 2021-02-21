package qiita

import (
	"reflect"
	"testing"
	"time"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
)

func TestSearchResult_GetCreatedDate(t *testing.T) {
	now := time.Now()
	type fields struct {
		RenderedBody   string
		Body           string
		Coediting      *bool
		CommentsCount  int
		CreatedAt      time.Time
		Group          string
		ID             string
		LikesCount     int
		Private        bool
		ReactionsCount int
		Tags           Tags
		Title          string
		UpdatedAt      time.Time
		URL            string
		User           User
		PageViewsCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系",
			fields: fields{
				CreatedAt: now,
			},
			want: now.Format(common.DateFormat4Show),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				RenderedBody:   tt.fields.RenderedBody,
				Body:           tt.fields.Body,
				Coediting:      tt.fields.Coediting,
				CommentsCount:  tt.fields.CommentsCount,
				CreatedAt:      tt.fields.CreatedAt,
				Group:          tt.fields.Group,
				ID:             tt.fields.ID,
				LikesCount:     tt.fields.LikesCount,
				Private:        tt.fields.Private,
				ReactionsCount: tt.fields.ReactionsCount,
				Tags:           tt.fields.Tags,
				Title:          tt.fields.Title,
				UpdatedAt:      tt.fields.UpdatedAt,
				URL:            tt.fields.URL,
				User:           tt.fields.User,
				PageViewsCount: tt.fields.PageViewsCount,
			}
			if got := r.GetCreatedDate(); got != tt.want {
				t.Errorf("GetCreatedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_GetProfileImageURL(t *testing.T) {
	type fields struct {
		RenderedBody   string
		Body           string
		Coediting      *bool
		CommentsCount  int
		CreatedAt      time.Time
		Group          string
		ID             string
		LikesCount     int
		Private        bool
		ReactionsCount int
		Tags           Tags
		Title          string
		UpdatedAt      time.Time
		URL            string
		User           User
		PageViewsCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系",
			fields: fields{
				User: User{
					ProfileImageURL: "https://hoge.com/a.png",
				},
			},
			want: "https://hoge.com/a.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				RenderedBody:   tt.fields.RenderedBody,
				Body:           tt.fields.Body,
				Coediting:      tt.fields.Coediting,
				CommentsCount:  tt.fields.CommentsCount,
				CreatedAt:      tt.fields.CreatedAt,
				Group:          tt.fields.Group,
				ID:             tt.fields.ID,
				LikesCount:     tt.fields.LikesCount,
				Private:        tt.fields.Private,
				ReactionsCount: tt.fields.ReactionsCount,
				Tags:           tt.fields.Tags,
				Title:          tt.fields.Title,
				UpdatedAt:      tt.fields.UpdatedAt,
				URL:            tt.fields.URL,
				User:           tt.fields.User,
				PageViewsCount: tt.fields.PageViewsCount,
			}
			if got := r.GetProfileImageURL(); got != tt.want {
				t.Errorf("GetProfileImageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_GetTags(t *testing.T) {
	type fields struct {
		RenderedBody   string
		Body           string
		Coediting      *bool
		CommentsCount  int
		CreatedAt      time.Time
		Group          string
		ID             string
		LikesCount     int
		Private        bool
		ReactionsCount int
		Tags           Tags
		Title          string
		UpdatedAt      time.Time
		URL            string
		User           User
		PageViewsCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "タグがある",
			fields: fields{
				Tags: []Tag{
					{
						Name:     "タグ1",
						Versions: []string{},
					},
					{
						Name:     "タグ2",
						Versions: []string{},
					},
					{
						Name:     "タグ3",
						Versions: []string{},
					},
				},
			},
			want: []string{
				"タグ1",
				"タグ2",
				"タグ3",
			},
		},
		{
			name: "タグがない",
			fields: fields{
				Tags: []Tag{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				RenderedBody:   tt.fields.RenderedBody,
				Body:           tt.fields.Body,
				Coediting:      tt.fields.Coediting,
				CommentsCount:  tt.fields.CommentsCount,
				CreatedAt:      tt.fields.CreatedAt,
				Group:          tt.fields.Group,
				ID:             tt.fields.ID,
				LikesCount:     tt.fields.LikesCount,
				Private:        tt.fields.Private,
				ReactionsCount: tt.fields.ReactionsCount,
				Tags:           tt.fields.Tags,
				Title:          tt.fields.Title,
				UpdatedAt:      tt.fields.UpdatedAt,
				URL:            tt.fields.URL,
				User:           tt.fields.User,
				PageViewsCount: tt.fields.PageViewsCount,
			}
			if got := r.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
