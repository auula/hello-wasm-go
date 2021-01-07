// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/1/7 - 4:26 下午 - UTC/GMT+08:00

package main

import (
	"fmt"
	"syscall/js"
)

// 通过Go代码操作页面dom节点
var (
	document = js.Global().Get("document")
	nameEle  = document.Call("getElementById", "name")
	helloEle = document.Call("getElementById", "hello")
	btnEle   = js.Global().Get("btn")
)

func sayHello(this js.Value, args []js.Value) interface{} {
	name := nameEle.Get("value").String()
	if len(name) == 0 {
		name = "github.com/higker/hello-wasm-go"
	}
	str := fmt.Sprintf("Hello,%s", name)
	helloEle.Set("innerHTML", js.ValueOf(str))
	return nil
}
func main() {
	done := make(chan int, 0)
	btnEle.Call("addEventListener", "click", js.FuncOf(sayHello))
	<-done
}
