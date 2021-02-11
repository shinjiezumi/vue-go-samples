package client

type Client interface {
	Init()
	Search(keyword string, count, page string)
}
