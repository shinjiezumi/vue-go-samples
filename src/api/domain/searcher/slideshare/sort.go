package slideshare

import "sort"

func Sort(res SearchResults) SearchResults {
	sort.SliceStable(res, func(i, j int) bool {
		// 更新日時の降順
		ret, _ := sortUpdatedDescFunc(res, i, j)

		return ret
	})

	return res
}

func sortUpdatedDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	iUpdated := res[i].GetUpdated()
	jUpdated := res[j].GetUpdated()

	if iUpdated.Equal(jUpdated) {
		return false, false
	}

	return iUpdated.After(jUpdated), true
}
