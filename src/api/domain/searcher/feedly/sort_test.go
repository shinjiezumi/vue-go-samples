package feedly

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
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
						Score: 100,
					},
					{
						Score: 102,
					},
					{
						Score: 101,
					},
					{
						Score: 103,
					},
				},
			},
			want: []SearchResult{
				{
					Score: 103,
				},
				{
					Score: 102,
				},
				{
					Score: 101,
				},
				{
					Score: 100,
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

func Test_sortScoreDescFunc(t *testing.T) {
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
			name: "スコアが同じ",
			args: args{
				res: []SearchResult{
					{
						Score: 100,
					},
					{
						Score: 100,
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
						Score: 101,
					},
					{
						Score: 100,
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
						Score: 100,
					},
					{
						Score: 101,
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
			gotRet, gotSorted := sortScoreDescFunc(tt.args.res, tt.args.i, tt.args.j)
			if gotRet != tt.wantRet {
				t.Errorf("sortScoreDescFunc() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
			if gotSorted != tt.wantSorted {
				t.Errorf("sortScoreDescFunc() gotSorted = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}
