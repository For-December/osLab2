package algorithm

type PageReplacementAlgorithm interface {

	// AccessPage 访问页面
	AccessPage(pageNumber int)

	// ReplacePage 找到要被替换的虚拟页号
	ReplacePage() int
}
