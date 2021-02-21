package slideshare

import (
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	now := time.Now()
	type args struct {
		res SearchResults
	}
	tests := []struct {
		name string
		args args
		want SearchResults
	}{
		{
			name: "正常系",
			args: args{
				res: []SearchResult{
					{
						ID:       1,
						Download: 10,
						Updated:  now.Add(-1 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
					{
						ID:       2,
						Download: 5,
						Updated:  now.Add(-2 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
					{
						ID:       3,
						Download: 15,
						Updated:  now.Add(-4 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
					{
						ID:       4,
						Download: 15,
						Updated:  now.Add(-5 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
					{
						ID:       5,
						Download: 15,
						Updated:  now.Add(-3 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
				},
			},
			want: []SearchResult{
				{
					ID:       5,
					Download: 15,
					Updated:  now.Add(-3 * time.Second).Format("2006-01-02 15:04:05 UTC"),
				},
				{
					ID:       3,
					Download: 15,
					Updated:  now.Add(-4 * time.Second).Format("2006-01-02 15:04:05 UTC"),
				},
				{
					ID:       4,
					Download: 15,
					Updated:  now.Add(-5 * time.Second).Format("2006-01-02 15:04:05 UTC"),
				},
				{
					ID:       1,
					Download: 10,
					Updated:  now.Add(-1 * time.Second).Format("2006-01-02 15:04:05 UTC"),
				},
				{
					ID:       2,
					Download: 5,
					Updated:  now.Add(-2 * time.Second).Format("2006-01-02 15:04:05 UTC"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDLCountDescFunc(t *testing.T) {
	type args struct {
		res SearchResults
		i   int
		j   int
	}
	tests := []struct {
		name       string
		args       args
		wantRet    bool
		wantSorted bool
	}{
		{
			name: "ダウンロード数が同じ",
			args: args{
				res: []SearchResult{
					{
						Download: 100,
					},
					{
						Download: 100,
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    false,
			wantSorted: false,
		},
		{
			name: "i > j",
			args: args{
				res: []SearchResult{
					{
						Download: 101,
					},
					{
						Download: 100,
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    true,
			wantSorted: true,
		},
		{
			name: "i > j",
			args: args{
				res: []SearchResult{
					{
						Download: 100,
					},
					{
						Download: 101,
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    false,
			wantSorted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, gotSorted := sortDLCountDescFunc(tt.args.res, tt.args.i, tt.args.j)
			if gotRet != tt.wantRet {
				t.Errorf("sortDLCountDescFunc() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
			if gotSorted != tt.wantSorted {
				t.Errorf("sortDLCountDescFunc() gotSorted = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}

func Test_sortUpdatedDescFunc(t *testing.T) {
	now := time.Now()
	type args struct {
		res SearchResults
		i   int
		j   int
	}
	tests := []struct {
		name       string
		args       args
		wantRet    bool
		wantSorted bool
	}{
		{
			name: "更新日時が同じ",
			args: args{
				res: []SearchResult{
					{
						Updated: now.Format("2006-01-02 15:04:05 UTC"),
					},
					{
						Updated: now.Format("2006-01-02 15:04:05 UTC"),
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    false,
			wantSorted: false,
		},
		{
			name: "i > j",
			args: args{
				res: []SearchResult{
					{
						Updated: now.Add(1 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
					{
						Updated: now.Format("2006-01-02 15:04:05 UTC"),
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    true,
			wantSorted: true,
		},
		{
			name: "i > j",
			args: args{
				res: []SearchResult{
					{
						Updated: now.Format("2006-01-02 15:04:05 UTC"),
					},
					{
						Updated: now.Add(1 * time.Second).Format("2006-01-02 15:04:05 UTC"),
					},
				},
				i: 0,
				j: 1,
			},
			wantRet:    false,
			wantSorted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, gotSorted := sortUpdatedDescFunc(tt.args.res, tt.args.i, tt.args.j)
			if gotRet != tt.wantRet {
				t.Errorf("sortUpdatedDescFunc() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
			if gotSorted != tt.wantSorted {
				t.Errorf("sortUpdatedDescFunc() gotSorted = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}
