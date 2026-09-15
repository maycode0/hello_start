package main

import (
	"fmt"
	"rsc.io/quote" // 外部模块里的 quote 包
)

func main() {
	fmt.Println(quote.Go())
}
