package qiita

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
						ID:         "1",
						LikesCount: 10,
						UpdatedAt:  now.Add(-1 * time.Second),
					},
					{
						ID:         "2",
						LikesCount: 5,
						UpdatedAt:  now.Add(-2 * time.Second),
					},
					{
						ID:         "3",
						LikesCount: 15,
						UpdatedAt:  now.Add(-4 * time.Second),
					},
					{
						ID:         "4",
						LikesCount: 15,
						UpdatedAt:  now.Add(-5 * time.Second),
					},
					{
						ID:         "5",
						LikesCount: 15,
						UpdatedAt:  now.Add(-3 * time.Second),
					},
				},
			},
			want: []SearchResult{
				{
					ID:         "5",
					LikesCount: 15,
					UpdatedAt:  now.Add(-3 * time.Second),
				},
				{
					ID:         "3",
					LikesCount: 15,
					UpdatedAt:  now.Add(-4 * time.Second),
				},
				{
					ID:         "4",
					LikesCount: 15,
					UpdatedAt:  now.Add(-5 * time.Second),
				},
				{
					ID:         "1",
					LikesCount: 10,
					UpdatedAt:  now.Add(-1 * time.Second),
				},
				{
					ID:         "2",
					LikesCount: 5,
					UpdatedAt:  now.Add(-2 * time.Second),
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

func Test_sortLGTMDescFunc(t *testing.T) {
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
			name: "LGTM数が同じ",
			args: args{
				res: []SearchResult{
					{
						LikesCount: 100,
					},
					{
						LikesCount: 100,
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
						LikesCount: 101,
					},
					{
						LikesCount: 100,
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
						LikesCount: 100,
					},
					{
						LikesCount: 101,
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
			gotRet, gotSorted := sortLGTMDescFunc(tt.args.res, tt.args.i, tt.args.j)
			if gotRet != tt.wantRet {
				t.Errorf("sortLGTMDescFunc() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
			if gotSorted != tt.wantSorted {
				t.Errorf("sortLGTMDescFunc() gotSorted = %v, want %v", gotSorted, tt.wantSorted)
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
						UpdatedAt: now,
					},
					{
						UpdatedAt: now,
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
						UpdatedAt: now.Add(1 * time.Second),
					},
					{
						UpdatedAt: now,
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
						UpdatedAt: now,
					},
					{
						UpdatedAt: now.Add(1 * time.Second),
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
