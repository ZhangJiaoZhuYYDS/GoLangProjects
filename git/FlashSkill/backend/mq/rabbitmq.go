// @Author zhangjiaozhu 2023/10/15 22:21:00
package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)
//ur1格式  amqp://账号：密码rabbitmq服务器地址：端口号/vhost

const MQURL = "amqp://demo:demo@192.168.31.5:5672/"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	// 对列名
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	//连接信息
	Mqurl string
}

func NewRabbitMQ(qn string, ec string,key string) *RabbitMQ {
	return &RabbitMQ{QueueName:qn,Exchange: ec,Key: key,Mqurl: MQURL}
}

func (r *RabbitMQ) Destroy()  {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string){
	if err != nil {
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}
//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName,"","")
	var err error
	//获取connection
	rabbitmq.conn,err  = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err,"failed to connect !")
	//获取channel
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err,"failed to open a channel")
	return rabbitmq
	}
func (r *RabbitMQ) PublishSimple(message string)  {
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		// 队列名
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否排他
		false,
		// 是否阻塞
		false,
		// 其他属性
		nil)
	if err != nil {
		log.Println(err)
	}
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true,根据exhange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange.发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
}

func( r *RabbitMQ) ConsumeSimple()  {
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		// 队列名
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否排他
		false,
		// 是否阻塞
		false,
		// 其他属性
		nil)
	if err != nil {
		log.Println(err)
	}
	msgs , err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
	}
	forever :=  make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			fmt.Println(d.Body)
		}
	}()
	<- forever
}
	