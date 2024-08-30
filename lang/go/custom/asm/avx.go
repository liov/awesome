package main

import (
	"fmt"
	"test/custom/asm/asm"
)

func main() {

	a := []int16{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	b := []int16{5, 15, 25, 35, 45, 55, 65, 75, 85, 95}

	// 调用汇编实现的 AVX2 函数
	result := asm.Avx2_ssd_int16(a, b, len(a))

	// 打印部分结果
	fmt.Println("Result (first 10 values):", result)
}
