package models

//// Frame 表示物理内存中的一个块
//type Frame struct {
//	FrameNumber int // 块号
//	PageNumber  int // 当前存储的页号，如果为空则为-1
//}

// Memory 表示物理内存
type Memory struct {
	FrameCount  int
	Frames      []int       // 记录每个帧当前分配的页面
	PageToFrame map[int]int // 页号到帧号的映射
}

func NewMemory(frameCount int) *Memory {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1 // -1表示帧为空
	}
	return &Memory{
		FrameCount:  frameCount,
		Frames:      frames,
		PageToFrame: make(map[int]int)}
}
