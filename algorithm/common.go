package algorithm

type PageReplacementAlgorithm interface {
	AccessPage(pageNumber int)
	ReplacePage() int
}
