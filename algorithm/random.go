package algorithm

import (
	"math/rand"
)

type RandomPageReplacement struct {
	Pages []int
}

func NewRandomPageReplacement(frameCount int) *RandomPageReplacement {
	return &RandomPageReplacement{Pages: make([]int, frameCount)}
}

func (random *RandomPageReplacement) AccessPage(pageNumber int) {
	for i, page := range random.Pages {
		if page == -1 {
			random.Pages[i] = pageNumber
			return
		}
	}
}

func (random *RandomPageReplacement) ReplacePage() int {

	// 随机选择一个页面进行替换，范围是 [0, len(random.Pages))
	pageIndex := rand.Intn(len(random.Pages))
	pageToReplace := random.Pages[pageIndex]
	random.Pages[pageIndex] = -1
	return pageToReplace
}
