package models

import (
	"cmp"
	"fmt"
	"osLab2/algorithm"
	"osLab2/utls/logger"
	"sort"
)

// VirtualMemoryManager 模拟操作系统的虚拟内存管理模块
type VirtualMemoryManager struct {
	memory    *Memory                            // 物理内存
	processes map[int]*Process                   // 所有的进程
	algorithm algorithm.PageReplacementAlgorithm // 页面替换算法接口
}

// NewVirtualMemoryManager 创建一个新的虚拟内存管理器, frameCount表示物理内存的帧数
func NewVirtualMemoryManager(
	frameCount int,
	algorithm algorithm.PageReplacementAlgorithm) *VirtualMemoryManager {
	return &VirtualMemoryManager{
		memory:    NewMemory(frameCount),
		processes: make(map[int]*Process),
		algorithm: algorithm,
	}
}

// AddProcess 添加一个进程到虚拟内存管理器(一般是刚创建的新进程)
func (vmm *VirtualMemoryManager) AddProcess(process *Process) {
	vmm.processes[process.PID] = process
}

// AccessAddress 访问一个进程的某个虚拟页号，返回物理页号和是否成功
func (vmm *VirtualMemoryManager) AccessAddress(processID, pageNumber int) (int, bool) {
	process, exists := vmm.processes[processID]
	if !exists {
		fmt.Printf("请求的进程 %d 不存在！\n", processID)
		return -1, false
	}

	// 内置二分查找
	index, exists := sort.Find(process.PageTable.Size, func(i int) int {
		return cmp.Compare(pageNumber,
			process.PageTable.Pages[i].PageNumber)
	})

	// 页表中不存在该虚拟页号
	if !exists {
		fmt.Println("虚拟页号不存在，out of index error！")
		return -1, false
	}

	// 获取页表中该虚拟页号的页元素
	page := process.PageTable.Pages[index]

	if process.PageTable.Pages[index].IsInMemory {
		logger.InfoF("进程 %d 的页号 %d 已经映射到了内存帧 %d，页表命中! ",
			processID, pageNumber, process.PageTable.Pages[index].FrameNumber)
		return page.FrameNumber, true
	}

	logger.ErrorF("进程 %d 的页号 %d 触发缺页错误", processID, pageNumber)
	return vmm.handlePageFault(process, page)
}

// handlePageFault 处理缺页错误，返回物理页号和是否成功
func (vmm *VirtualMemoryManager) handlePageFault(process *Process, page *PageTableEntry) (int, bool) {

	// 查找空闲帧
	frameNumber := vmm.findFreeFrame()

	// 如果没有空闲帧，执行页面替换算法
	if frameNumber == -1 {
		pageToReplace := vmm.algorithm.ReplacePage()

		// 如果没有页面需要替换
		if pageToReplace == -1 {
			fmt.Println("没有页面需要替换")
			return -1, false
		}

		// 找到需要替换的页面对应的帧号
		frameNumber = vmm.memory.PageToFrame[pageToReplace]
		fmt.Printf("Replacing page %d from frame %d\n", pageToReplace, frameNumber)
		for _, proc := range vmm.processes {

			// 内置二分查找
			index, exists := sort.Find(process.PageTable.Size, func(i int) int {
				return cmp.Compare(pageToReplace,
					process.PageTable.Pages[i].PageNumber)
			})

			if exists {
				proc.PageTable.Pages[index].IsInMemory = false
				proc.PageTable.Pages[index].FrameNumber = -1
				break
			}
		}
	}

	// 加载页面到帧
	vmm.loadPageIntoFrame(process, page, frameNumber)
	return frameNumber, true
}

// findFreeFrame 查找空闲帧
func (vmm *VirtualMemoryManager) findFreeFrame() int {
	for i, frame := range vmm.memory.Frames {
		if frame == -1 {
			return i
		}
	}
	return -1
}

// loadPageIntoFrame 将页面加载到帧
func (vmm *VirtualMemoryManager) loadPageIntoFrame(process *Process, page *PageTableEntry, frameNumber int) {
	vmm.memory.Frames[frameNumber] = page.PageNumber
	vmm.memory.PageToFrame[page.PageNumber] = frameNumber
	page.IsInMemory = true
	page.FrameNumber = frameNumber
	vmm.algorithm.AccessPage(page.PageNumber)
	fmt.Printf("Loaded page %d of process %d into frame %d\n", page.PageNumber, process.PID, frameNumber)
}
