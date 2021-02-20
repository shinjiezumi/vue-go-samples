package slideshare

import "sort"

func Sort(res SearchResults) SearchResults {
	sort.SliceStable(res, func(i, j int) bool {
		// ダウンロード数の降順
		if ret, sorted := sortDLCountDescFunc(res, i, j); sorted {
			return ret
		}

		// 更新日時の降順
		ret, _ := sortUpdatedDescFunc(res, i, j)

		return ret
	})

	return res
}

func sortDLCountDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	if res[i].Download == res[j].Download {
		return false, false
	}

	return res[i].Download > res[j].Download, true
}

func sortUpdatedDescFunc(res SearchResults, i, j int) (ret, sorted bool) {
	iUpdated := res[i].GetUpdated()
	jUpdated := res[j].GetUpdated()

	if iUpdated.Equal(jUpdated) {
		return false, false
	}

	return iUpdated.After(jUpdated), true
}
