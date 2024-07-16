package main

import (
	"fmt"
	"osLab2/algorithm"
	"osLab2/models"
)

// 测试案例：https://m.ofweek.com/ai/2021-04/ART-201721-11000-30495592.html
func main() {
	// 初始化虚拟内存管理器，设置物理内存帧数和页面置换算法
	vmm := models.NewVirtualMemoryManager(3, algorithm.NewFIFOPageReplacement())

	// 创建进程并添加到虚拟内存管理器中
	// 真实物理帧只有3个，而虚拟页号有5个（5个虚拟页共享3个物理帧），所以会发生缺页错误和页面置换情况
	process1 := models.NewProcess(1, 5)
	vmm.AddProcess(process1)

	// 模拟访问地址{0, 1, 2, 0, 3, 4, 1, 0, 2}

	// 测试案例来自王道：https://m.ofweek.com/ai/2021-04/ART-201721-11000-30495592.html
	pagesToAccess := []int{3, 2, 1, 0, 3, 2, 4, 3, 2, 1, 0, 4}

	for _, pageNumber := range pagesToAccess {
		frameNumber, success := vmm.AccessAddress(1, pageNumber)
		if success {
			fmt.Printf("Accessed page %d in frame %d\n", pageNumber, frameNumber)
		} else {
			fmt.Printf("Failed to access page %d\n", pageNumber)
		}
	}
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
