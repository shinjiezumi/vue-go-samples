package feedly

import (
	"reflect"
	"testing"
)

var dummyStr = "aaabbbccc"

func TestSearchResponseWrapper_GetSearchResults(t *testing.T) {
	type fields struct {
		Results SearchResults
	}
	tests := []struct {
		name   string
		fields fields
		want   SearchResults
	}{
		{
			name: "正常系",
			fields: fields{
				Results: []SearchResult{
					{
						FeedID: dummyStr,
					},
				},
			},
			want: []SearchResult{
				{
					FeedID: dummyStr,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &SearchResponseWrapper{
				Results: tt.fields.Results,
			}
			if got := w.GetSearchResults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSearchResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_GetDescription(t *testing.T) {
	type fields struct {
		FeedID        string
		Subscribers   int
		Score         float32
		Title         string
		Description   *string
		Website       *string
		LastUpdated   *int
		Velocity      *float32
		Language      *string
		Featured      *bool
		IconURL       *string
		VisualURL     *string
		CoverURL      *string
		Logo          *string
		ContentType   *string
		CoverColor    *string
		DeliciousTags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "サイト概要がnilでない",
			fields: fields{
				Description: &dummyStr,
			},
			want: dummyStr,
		},
		{
			name: "サイト概要がnil",
			fields: fields{
				Description: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				FeedID:        tt.fields.FeedID,
				Subscribers:   tt.fields.Subscribers,
				Score:         tt.fields.Score,
				Title:         tt.fields.Title,
				Description:   tt.fields.Description,
				Website:       tt.fields.Website,
				LastUpdated:   tt.fields.LastUpdated,
				Velocity:      tt.fields.Velocity,
				Language:      tt.fields.Language,
				Featured:      tt.fields.Featured,
				IconURL:       tt.fields.IconURL,
				VisualURL:     tt.fields.VisualURL,
				CoverURL:      tt.fields.CoverURL,
				Logo:          tt.fields.Logo,
				ContentType:   tt.fields.ContentType,
				CoverColor:    tt.fields.CoverColor,
				DeliciousTags: tt.fields.DeliciousTags,
			}
			if got := r.GetDescription(); got != tt.want {
				t.Errorf("GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_GetSiteImageURL1(t *testing.T) {
	type fields struct {
		FeedID        string
		Subscribers   int
		Score         float32
		Title         string
		Description   *string
		Website       *string
		LastUpdated   *int
		Velocity      *float32
		Language      *string
		Featured      *bool
		IconURL       *string
		VisualURL     *string
		CoverURL      *string
		Logo          *string
		ContentType   *string
		CoverColor    *string
		DeliciousTags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "サイト画像URLがnilでない",
			fields: fields{
				VisualURL: &dummyStr,
			},
			want: dummyStr,
		},
		{
			name: "サイト画像URLがnil",
			fields: fields{
				VisualURL: nil,
			},
			want: "https://placehold.jp/150x150.png?text=NO+IMAGE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				FeedID:        tt.fields.FeedID,
				Subscribers:   tt.fields.Subscribers,
				Score:         tt.fields.Score,
				Title:         tt.fields.Title,
				Description:   tt.fields.Description,
				Website:       tt.fields.Website,
				LastUpdated:   tt.fields.LastUpdated,
				Velocity:      tt.fields.Velocity,
				Language:      tt.fields.Language,
				Featured:      tt.fields.Featured,
				IconURL:       tt.fields.IconURL,
				VisualURL:     tt.fields.VisualURL,
				CoverURL:      tt.fields.CoverURL,
				Logo:          tt.fields.Logo,
				ContentType:   tt.fields.ContentType,
				CoverColor:    tt.fields.CoverColor,
				DeliciousTags: tt.fields.DeliciousTags,
			}
			if got := r.GetSiteImageURL(); got != tt.want {
				t.Errorf("GetSiteImageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_GetSiteURL(t *testing.T) {
	type fields struct {
		FeedID        string
		Subscribers   int
		Score         float32
		Title         string
		Description   *string
		Website       *string
		LastUpdated   *int
		Velocity      *float32
		Language      *string
		Featured      *bool
		IconURL       *string
		VisualURL     *string
		CoverURL      *string
		Logo          *string
		ContentType   *string
		CoverColor    *string
		DeliciousTags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "サイトURLがnilでない",
			fields: fields{
				Website: &dummyStr,
			},
			want: dummyStr,
		},
		{
			name: "サイトURLがnil",
			fields: fields{
				Website: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				FeedID:        tt.fields.FeedID,
				Subscribers:   tt.fields.Subscribers,
				Score:         tt.fields.Score,
				Title:         tt.fields.Title,
				Description:   tt.fields.Description,
				Website:       tt.fields.Website,
				LastUpdated:   tt.fields.LastUpdated,
				Velocity:      tt.fields.Velocity,
				Language:      tt.fields.Language,
				Featured:      tt.fields.Featured,
				IconURL:       tt.fields.IconURL,
				VisualURL:     tt.fields.VisualURL,
				CoverURL:      tt.fields.CoverURL,
				Logo:          tt.fields.Logo,
				ContentType:   tt.fields.ContentType,
				CoverColor:    tt.fields.CoverColor,
				DeliciousTags: tt.fields.DeliciousTags,
			}
			if got := r.GetSiteURL(); got != tt.want {
				t.Errorf("GetSiteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
