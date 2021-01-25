package searcher

type ApiDriver interface {
	Init()
	Search(keyword string, count, page string)
}
