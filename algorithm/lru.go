package algorithm

type LRUPageReplacement struct {
	Pages []int
}

func NewLRUPageReplacement() *LRUPageReplacement {
	return &LRUPageReplacement{Pages: []int{}}
}

func (lru *LRUPageReplacement) AccessPage(pageNumber int) {
	// 移除已经存在的页面
	for i, page := range lru.Pages {
		if page == pageNumber {
			lru.Pages = append(lru.Pages[:i], lru.Pages[i+1:]...)
			break
		}
	}
	lru.Pages = append(lru.Pages, pageNumber)
}

func (lru *LRUPageReplacement) ReplacePage() int {
	if len(lru.Pages) == 0 {
		return -1
	}
	pageToReplace := lru.Pages[0]
	lru.Pages = lru.Pages[1:]
	return pageToReplace
}
