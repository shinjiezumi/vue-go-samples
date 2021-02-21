package slideshare

import (
	"reflect"
	"testing"
	"time"
)

var dummyNum = 1

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
						ID: dummyNum,
					},
				},
			},
			want: []SearchResult{
				{
					ID: dummyNum,
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

func TestSearchResult_GetUpdated(t *testing.T) {
	d := "2021-02-21 00:00:00 UTC"
	type fields struct {
		ID                  int
		Title               string
		Description         string
		Username            string
		URL                 string
		ThumbnailURL        string
		ThumbnailSize       ThumbnailSize
		ThumbnailSmallURL   string
		ThumbnailXLargeURL  string
		ThumbnailXXLargeURL string
		Embed               string
		Created             string
		Updated             string
		Language            string
		Format              string
		Download            bool
		DownloadURL         string
		SecretKey           string
		SlideshowEmbedURL   string
		SlideshowType       int
		InContest           int
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "正常系",
			fields: fields{
				Updated: d,
			},
			want: func() time.Time {
				t, err := time.Parse("2006-01-02 15:04:05 UTC", d)
				if err != nil {
					panic(err)
				}
				return t
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResult{
				ID:                  tt.fields.ID,
				Title:               tt.fields.Title,
				Description:         tt.fields.Description,
				Username:            tt.fields.Username,
				URL:                 tt.fields.URL,
				ThumbnailURL:        tt.fields.ThumbnailURL,
				ThumbnailSize:       tt.fields.ThumbnailSize,
				ThumbnailSmallURL:   tt.fields.ThumbnailSmallURL,
				ThumbnailXLargeURL:  tt.fields.ThumbnailXLargeURL,
				ThumbnailXXLargeURL: tt.fields.ThumbnailXXLargeURL,
				Embed:               tt.fields.Embed,
				Created:             tt.fields.Created,
				Updated:             tt.fields.Updated,
				Language:            tt.fields.Language,
				Format:              tt.fields.Format,
				Download:            tt.fields.Download,
				DownloadURL:         tt.fields.DownloadURL,
				SecretKey:           tt.fields.SecretKey,
				SlideshowEmbedURL:   tt.fields.SlideshowEmbedURL,
				SlideshowType:       tt.fields.SlideshowType,
				InContest:           tt.fields.InContest,
			}
			if got := r.GetUpdated(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdated() = %v, want %v", got, tt.want)
			}
		})
	}
}
