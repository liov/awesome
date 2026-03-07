package main

import (
	"context"
	"fmt"
)

// 1. 定义一个嵌入 context.Context 的结构体
type myCtx struct {
	context.Context
}

func checkType(p context.Context) {
	// 尝试断言 p 的动态类型是否为 context.Context 接口
	_, ok := p.(context.Context)

	fmt.Printf("动态类型是: %T\n", p)
	fmt.Printf("断言 p.(context.Context) 的结果 ok = %v\n", ok)

	// 正确的做法：如果你想确认它是不是你的自定义类型
	_, ok2 := p.(myCtx)
	fmt.Printf("断言 p.(myCtx) 的结果 ok = %v\n", ok2)

	// 或者，如果你想获取底层的 context.Context (通常没必要，因为 p 本身就是)
	// 如果 myCtx 里还包了一层，你可能需要访问字段
	if m, isMy := p.(myCtx); isMy {
		fmt.Println("成功识别为 myCtx，可以直接调用 m.Context 的方法")
		_ = m.Context
	}
}

func main() {
	// 创建一个 myCtx 实例，嵌入 context.TODO()
	c := myCtx{
		Context: context.TODO(),
	}

	// 传给函数
	checkType(c)
}