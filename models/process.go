package models

import (
	"cmp"
	"errors"
	"osLab2/constraint"
	"sort"
)

// PageTableEntry 页表实体元素
type PageTableEntry struct {
	PageNumber  int // 页号
	FrameNumber int // 物理块号
	IsInMemory  bool
}

// PageTable 表示页表，维护页号与物理块号的映射关系
type PageTable struct {
	Pages []*PageTableEntry // 页号到页面的映射
	Size  int
}

type Process struct {
	PID       int
	PageTable PageTable // 页表
}

// NewProcess 创建一个新的进程
func NewProcess(id int, pageCount int) *Process {
	pageTable := PageTable{
		Pages: make([]*PageTableEntry, pageCount),
		Size:  pageCount,
	}

	// 初始化页表，所有页面均未加载到内存
	for i := 0; i < pageCount; i++ {
		pageTable.Pages[i] =
			&PageTableEntry{
				PageNumber:  i,
				FrameNumber: -1,
				IsInMemory:  false}
	}

	return &Process{PID: id, PageTable: pageTable}
}

// Translate 地址转换函数，将逻辑地址转换为物理地址
func (pt *PageTable) Translate(logicalAddress int) (int, error) {
	pageNumber := logicalAddress / constraint.PageSize
	offset := logicalAddress % constraint.PageSize

	// 内置二分查找
	index, exists := sort.Find(pt.Size, func(i int) int {
		return cmp.Compare(pageNumber,
			pt.Pages[i].PageNumber)
	})
	if !exists {
		return -1, errors.New("缺页中断，需执行页替换算法")
	}
	if pt.Pages[index].FrameNumber == -1 {
		return -1, errors.New("因首次加载页而产生缺页中断")
	}
	physicalAddress := pt.Pages[index].FrameNumber*constraint.PageSize + offset
	return physicalAddress, nil
}
