package models

import (
	"cmp"
	"fmt"
	"osLab2/algorithm"
	"sort"
)

type VirtualMemoryManager struct {
	memory    *Memory
	processes map[int]*Process
	algorithm algorithm.PageReplacementAlgorithm
}

func NewVirtualMemoryManager(
	frameCount int,
	algorithm algorithm.PageReplacementAlgorithm) *VirtualMemoryManager {
	return &VirtualMemoryManager{
		memory:    NewMemory(frameCount),
		processes: make(map[int]*Process),
		algorithm: algorithm,
	}
}

func (vmm *VirtualMemoryManager) AddProcess(process *Process) {
	vmm.processes[process.PID] = process
}

func (vmm *VirtualMemoryManager) AccessAddress(processID, pageNumber int) (int, bool) {
	process, exists := vmm.processes[processID]
	if !exists {
		fmt.Println("Process not found")
		return -1, false
	}
	// 内置二分查找
	index, exists := sort.Find(process.PageTable.Size, func(i int) int {
		return cmp.Compare(pageNumber,
			process.PageTable.Pages[i].PageNumber)
	})

	if !exists {
		fmt.Println("Page not found in process page table")
		return -1, false
	}

	page := process.PageTable.Pages[index]

	if process.PageTable.Pages[index].IsInMemory {
		fmt.Printf("Page %d in process %d is already in memory at frame %d\n",
			pageNumber, processID, process.PageTable.Pages[index].FrameNumber)
		return page.FrameNumber, true
	}

	fmt.Printf("Page %d in process %d caused a page fault\n", pageNumber, processID)
	return vmm.handlePageFault(process, page)
}

func (vmm *VirtualMemoryManager) handlePageFault(process *Process, page *PageTableEntry) (int, bool) {
	return 1, false
	//frameNumber := vmm.findFreeFrame()
	//if frameNumber == -1 {
	//	pageToReplace := vmm.algorithm.ReplacePage()
	//	if pageToReplace == -1 {
	//		fmt.Println("No page to replace")
	//		return -1, false
	//	}
	//	frameNumber = vmm.memory.PageToFrame[pageToReplace]
	//	fmt.Printf("Replacing page %d from frame %d\n", pageToReplace, frameNumber)
	//	for _, proc := range vmm.processes {
	//		if proc.PageTable[pageToReplace] != nil {
	//			proc.PageTable[pageToReplace].IsInMemory = false
	//			proc.PageTable[pageToReplace].FrameNumber = -1
	//			break
	//		}
	//	}
	//}
	//vmm.loadPageIntoFrame(process, page, frameNumber)
	//return frameNumber, true
}

func (vmm *VirtualMemoryManager) findFreeFrame() int {
	for i, frame := range vmm.memory.Frames {
		if frame == -1 {
			return i
		}
	}
	return -1
}

func (vmm *VirtualMemoryManager) loadPageIntoFrame(process *Process, page *PageTableEntry, frameNumber int) {
	//vmm.memory.Frames[frameNumber] = page.PageNumber
	//vmm.memory.PageToFrame[page.PageNumber] = frameNumber
	//page.IsInMemory = true
	//page.FrameNumber = frameNumber
	//vmm.algorithm.AccessPage(page.PageNumber)
	//fmt.Printf("Loaded page %d of process %d into frame %d\n", page.PageNumber, process.ID, frameNumber)
}
