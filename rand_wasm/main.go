package main

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// 将WebAssembly模块读取为字节
	bytes, _ := wasm.ReadBytes("tests.wasm")

	// 实例化WebAssembly模块
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// 从WebAssembly实例获取`sum`导出的函数。
	sum := instance.Exports["sum"]

	// 用Go标准值调用导出的函数。WebAssembly
	// 推断类型，并自动转换值
	result, _ := sum(11, 11)

	// 结果
	fmt.Println(result) // 22!
}
