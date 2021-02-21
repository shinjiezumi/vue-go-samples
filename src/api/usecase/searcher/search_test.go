package searcher

import (
	"reflect"
	"testing"
)

func Test_searchUseCase_parseQuery(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "クエリ1個",
			args: args{
				q: "hoge",
			},
			want: []string{
				"hoge",
			},
		},
		{
			name: "クエリ複数",
			args: args{
				q: "hoge,fuga,hogehoge",
			},
			want: []string{
				"hoge",
				"fuga",
				"hogehoge",
			},
		},
		{
			name: "クエリ重複",
			args: args{
				q: "hoge,fuga,hoge,fuga",
			},
			want: []string{
				"hoge",
				"fuga",
			},
		},
		{
			name: "クエリなし",
			args: args{
				q: "",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &searchUseCase{}
			if got := s.parseQuery(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
