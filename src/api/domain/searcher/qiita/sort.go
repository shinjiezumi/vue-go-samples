package qiita

import "sort"

func Sort(res SearchResults) SearchResults {
	sort.SliceStable(res, func(i, j int) bool {
		// LGTM数の降順
		if ret, sorted := sortLGTMDescFunc(res, i, j); sorted {
			return ret
		}

		// 更新日時の降順
		ret, _ := sortUpdatedDescFunc(res, i, j)

		return ret
	})

	return res
}

func sortLGTMDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	if res[i].LikesCount == res[j].LikesCount {
		return false, false
	}

	return res[i].LikesCount > res[j].LikesCount, true
}

func sortUpdatedDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	if res[i].UpdatedAt.Equal(res[j].UpdatedAt) {
		return false, false
	}

	return res[i].UpdatedAt.After(res[j].UpdatedAt), true
}
