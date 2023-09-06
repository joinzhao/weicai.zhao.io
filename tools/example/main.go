package main

import (
	"fmt"
	"weicai.zhao.io/tools"
)

func main() {
	fmt.Println(tools.IsEmpty(1))
	fmt.Println(tools.IsEmpty(1.1))
	fmt.Println(tools.IsEmpty(-10))
	fmt.Println(tools.IsEmpty(-10.1))
	fmt.Println(tools.IsEmpty("1231"))
	fmt.Println(tools.IsEmpty(""))
	fmt.Println(tools.IsEmpty(0))
}
