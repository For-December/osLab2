package algorithm

type FIFOPageReplacement struct {
	Queue []int
}

func NewFIFOPageReplacement() *FIFOPageReplacement {
	return &FIFOPageReplacement{Queue: []int{}}
}

func (fifo *FIFOPageReplacement) AccessPage(pageNumber int) {
	fifo.Queue = append(fifo.Queue, pageNumber)
}

func (fifo *FIFOPageReplacement) ReplacePage() int {
	if len(fifo.Queue) == 0 {
		return -1
	}
	pageToReplace := fifo.Queue[0]
	fifo.Queue = fifo.Queue[1:]
	return pageToReplace
}
