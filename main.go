package main

import (
	"fmt"
	"osLab2/constraint"
	"osLab2/models"
)

func main() {
	// 示例：创建一个页表和物理内存
	pageTable := models.NewPageTable(3)
	memory := models.NewMemory(4)

	// 添加页面到页表
	pageTable.Pages[0] = &models.PageTableEntry{PageNumber: 0, FrameNumber: -1}
	pageTable.Pages[1] = &models.PageTableEntry{PageNumber: 1, FrameNumber: -1}

	// 分配物理块给页面
	frameNumber, err := memory.AllocateFrame(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		pageTable.Pages[0].FrameNumber = frameNumber
	}

	frameNumber, err = memory.AllocateFrame(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		pageTable.Pages[1].FrameNumber = frameNumber
	}

	// 逻辑地址转换为物理地址
	logicalAddress := 2048 // 页号为2，偏移量为0
	physicalAddress, err := pageTable.Translate(logicalAddress)
	if err != nil {
		fmt.Println("Translate Error:", err)
		// 处理缺页中断
		pageNumber := logicalAddress / constraint.PageSize
		err = models.HandlePageFault(pageTable, memory, pageNumber)
		if err != nil {
			fmt.Println("Handle Page Fault Error:", err)
		} else {
			// 再次尝试地址转换
			physicalAddress, err = pageTable.Translate(logicalAddress)
			if err != nil {
				fmt.Println("Translate Error after Page Fault:", err)
			} else {
				fmt.Printf("Logical Address: %d -> Physical Address: %d\n", logicalAddress, physicalAddress)
			}
		}
	} else {
		fmt.Printf("Logical Address: %d -> Physical Address: %d\n", logicalAddress, physicalAddress)
	}

	pageTableStr := "当前页表 =>\n"
	for _, entry := range pageTable.Pages {
		pageTableStr += fmt.Sprintf("页号：%d，块号：%d\n", entry.PageNumber, entry.FrameNumber)
	}
	fmt.Println(pageTableStr)
	//fmt.Printf("Page Table: %+v\n", PrettyPrint(pageTable.Pages))
	//fmt.Printf("Memory: %+v\n", PrettyPrint(memory.Frames))
}

//func PrettyPrint(v interface{}) string {
//	b, err := json.Marshal(v)
//	if err != nil {
//		fmt.Println(v)
//		return ""
//	}
//
//	var out bytes.Buffer
//	err = json.Indent(&out, b, "", "  ")
//	if err != nil {
//		fmt.Println(v)
//		return ""
//	}
//
//	return out.String()
//}
