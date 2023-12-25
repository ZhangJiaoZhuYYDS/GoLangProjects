// @Author zhangjiaozhu 2023/11/25 9:20:00
package main

import (
	"car/conf"
	"car/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
