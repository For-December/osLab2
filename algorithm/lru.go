package algorithm

type LRUPageReplacement struct {
	Pages []int
}

func NewLRUPageReplacement() *LRUPageReplacement {
	return &LRUPageReplacement{Pages: []int{}}
}

func (lru *LRUPageReplacement) AccessPage(pageNumber int) {
	// 将最近访问的页面移动到最后
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

	// 移除最近最少使用的页面，即第一个页面
	pageToReplace := lru.Pages[0]
	lru.Pages = lru.Pages[1:]
	return pageToReplace
}
