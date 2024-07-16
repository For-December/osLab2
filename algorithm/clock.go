package algorithm

type CLOCKPageReplacement struct {
	Clock []bool
	Hand  int
	Pages []int
}

func NewCLOCKPageReplacement(frameCount int) *CLOCKPageReplacement {
	return &CLOCKPageReplacement{
		Clock: make([]bool, frameCount),
		Hand:  0,
		Pages: make([]int, frameCount),
	}
}

func (clock *CLOCKPageReplacement) AccessPage(pageNumber int) {
	for i, page := range clock.Pages {
		if page == pageNumber {
			clock.Clock[i] = true
			return
		}
	}
	for i := range clock.Pages {
		if clock.Pages[i] == -1 {
			clock.Pages[i] = pageNumber
			clock.Clock[i] = true
			return
		}
	}
}

func (clock *CLOCKPageReplacement) ReplacePage() int {
	for {
		if !clock.Clock[clock.Hand] {
			pageToReplace := clock.Pages[clock.Hand]
			clock.Pages[clock.Hand] = -1
			clock.Hand = (clock.Hand + 1) % len(clock.Clock)
			return pageToReplace
		}
		clock.Clock[clock.Hand] = false
		clock.Hand = (clock.Hand + 1) % len(clock.Clock)
	}
}
