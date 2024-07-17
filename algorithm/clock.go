package algorithm

type CLOCKPageReplacement struct {
	Clock []bool
	Ptr   int // 时钟指针，指向下一个要被替换的页面
	Pages []int
}

func NewCLOCKPageReplacement(frameCount int) *CLOCKPageReplacement {
	return &CLOCKPageReplacement{
		Clock: make([]bool, frameCount),
		Ptr:   0,
		Pages: make([]int, frameCount),
	}
}

func (clock *CLOCKPageReplacement) AccessPage(pageNumber int) {
	for i, page := range clock.Pages {
		if page == pageNumber {
			clock.Clock[i] = true // 标记为已访问
			return
		}
	}
	for i := range clock.Pages {
		if clock.Pages[i] == -1 {
			clock.Pages[i] = pageNumber // 将页面添加到空闲位置
			clock.Clock[i] = true
			return
		}
	}
}

func (clock *CLOCKPageReplacement) ReplacePage() int {
	for {

		// 如果页面被访问过，则将其标记为未访问，并继续查找
		if !clock.Clock[clock.Ptr] {

			// 找到要被替换的页面
			pageToReplace := clock.Pages[clock.Ptr]
			clock.Pages[clock.Ptr] = -1                    // 标记为空闲
			clock.Ptr = (clock.Ptr + 1) % len(clock.Clock) // 循环移动指针
			return pageToReplace
		}

		// 如果页面未被访问过，则将其标记为未访问，并返回
		clock.Clock[clock.Ptr] = false
		clock.Ptr = (clock.Ptr + 1) % len(clock.Clock)
	}
}
