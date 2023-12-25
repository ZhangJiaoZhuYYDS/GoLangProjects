// @Author zhangjiaozhu 2023/6/8 13:56:00
package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("i", 100, "输入数字")
	n := flag.String("name", "wang", "输入名字")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*i, *n)
}
