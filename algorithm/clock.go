package algorithm

type CLOCKPageReplacement struct {
	Clock []bool // 时钟指针数组，用于标记页面是否被访问,0 1
	Ptr   int    // 时钟指针，指向下一个要被替换的页面
	Pages []int  // 页面数组，指针行走决定替换哪个页面
}

func NewCLOCKPageReplacement(frameCount int) *CLOCKPageReplacement {
	pages := make([]int, frameCount)

	// fixBugs: 修复了pages数组的初始化问题
	// pages数组初始化为-1，表示为空
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

	// 时钟包含所有物理帧，每个物理帧对应一个页面
	for i, page := range clock.Pages {
		if page == pageNumber {
			clock.Clock[i] = true // 标记为已访问
			return
		}
	}
	for i := range clock.Pages {
		if clock.Pages[i] == -1 {
			clock.Pages[i] = pageNumber // 将页号放入时钟中轮转
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
