// @Author zhangjiaozhu 2023/10/3 23:59:00
package main

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"log"
)

func main() {
	pusher := kq.NewPusher([]string{"192.168.31.5:9092"}, "topic-beyond-like")
	marshal, _ := json.Marshal("你好啊")
	if err := pusher.Push(string(marshal)); err != nil {
		log.Println(err)
		return
	}else {
		log.Println("success")
	}

}
