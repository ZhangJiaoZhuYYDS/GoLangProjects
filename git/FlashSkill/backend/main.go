// @Author zhangjiaozhu 2023/10/14 17:12:00
package main

import "FlashSkill/backend/mq"

func main() {
	//_, err := common.NewMysqlConn()
	//if err != nil {
	//	log.Fatalf("Error: %v\n", err)
	//}
	//
	//r := router.New()
	//r.Run(":8080")

	//Simple模式 发送者
	rabbitmq := mq.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("hello imooc!")
	////接收者
	rabbit := mq.NewRabbitMQSimple("imoocSimple")
	rabbit.ConsumeSimple()

	//msgs := make(chan bool)
	//
	//forever := make(chan bool)
	//// 启用协程处理消息
	//go func(c int) {
	//	for d := range msgs {
	//		// 实现我们要处理的逻辑函数
	//		fmt.Println(d,c)
	//		c++
	//	}
	//	close(forever) // 关闭 forever 通道以通知主程序退出
	//}(1)
	//
	//for i := 0; i < 1000; i++ {
	//	msgs <- true
	//}
	//// 阻塞等待主程序退出
	//<-forever
	//time.Sleep(10*time.Second)
}
