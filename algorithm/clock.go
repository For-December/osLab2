package algorithm

type CLOCKPageReplacement struct {
	Clock []bool
	Ptr   int // 时钟指针，指向下一个要被替换的页面
	Pages []int
}

func NewCLOCKPageReplacement(frameCount int) *CLOCKPageReplacement {
	pages := make([]int, frameCount)

	// fixBugs: 修复了pages数组的初始化问题
	for i := range pages {
		pages[i] = -1
	}
	return &CLOCKPageReplacement{
		Clock: make([]bool, frameCount),
		Ptr:   0,
		Pages: pages,
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
		// 如果是0，则替换页面
		if !clock.Clock[clock.Ptr] {
			pageToReplace := clock.Pages[clock.Ptr]
			clock.Pages[clock.Ptr] = -1                    // 标记为空闲
			clock.Ptr = (clock.Ptr + 1) % len(clock.Clock) // 循环移动指针
			return pageToReplace
		}

		// 如果是1，改为0，继续下一个
		clock.Clock[clock.Ptr] = false
		clock.Ptr = (clock.Ptr + 1) % len(clock.Clock)
	}
}
