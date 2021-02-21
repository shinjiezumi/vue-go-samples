package feedly

import "sort"

func Sort(res SearchResults) SearchResults {
	sort.SliceStable(res, func(i, j int) bool {
		// スコアの降順
		ret, _ := sortScoreDescFunc(res, i, j)

		return ret
	})

	return res
}

func sortScoreDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	if res[i].Score == res[j].Score {
		return false, false
	}

	return res[i].Score > res[j].Score, true
}
