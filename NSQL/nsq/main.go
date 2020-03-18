// nsq go语言开发的消息对列
//具体安装查看官方文档 此处使用docker镜像
//https: //nsq.io/overview/quick_start.html
/**
docker pull nsqio/nsq:v1.2.0
sudo docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq:v1.2.0 /nsqlookupd

sudo docker run -d --name nsqd   -p 4150:4150 -p 4151:4151 \
	nsqio/nsq:v1.2.0 /nsqd \
	--broadcast-address=192.168.118.129 \
	--lookupd-tcp-address=192.168.118.129:4160


sudo docker run -d --name nsqadmin  -p 4171:4171 \
    nsqio/nsq:v1.2.0 /nsqadmin \
    --lookupd-http-address=192.168.118.129:4161 
	
记得关闭防火墙 路由找不到
	
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	nsq "github.com/nsqio/go-nsq"
)


	var conf *nsq.Config

	var sw sync.WaitGroup

	var err error

	var rec =func(){
		if er:=recover();er != nil{
			fmt.Printf("发生panic： %v",er)
		}
	}
	type customerMessageHandler struct {}

//初始化
func init(){
	conf = nsq.NewConfig()
	
}


func main() {
	//生产者消费者 topic 
	sw.Add(4)
	go producer()
	go customer()
	go customer()
	go customer()
	sw.Wait()


}


func(c *customerMessageHandler)HandleMessage(message *nsq.Message) error{
	defer time.Sleep(1 * time.Second)
	if len(message.Body) == 0 {
		return nil
	} 
		fmt.Println("消费消息："+string(message.Body))	
		return nil
		defer
}

func(c *customerMessageHandler)pri(){
	
}


//消费者
func customer(){
	defer sw.Done()
	defer rec()
	//192.168.118.129:4161
	customer,err:= nsq.NewConsumer("topic","chan1",conf)
	if err != nil {
		panic(err)
	}
	 customer.AddHandler(&customerMessageHandler{})
	 err =customer.ConnectToNSQLookupds([]string{"192.168.118.129:4161"})
	 if err != nil {
		panic(err)
	}
	fmt.Println("连接NSQLookupd")
	time.Sleep(100 * time.Second)
	customer.Stop()
	fmt.Println("断开NSQLookupd")

}

//生产者
func producer(){
	defer sw.Done()
	defer rec()
	//生产者 连接 nsqd
	prod,err:= nsq.NewProducer("192.168.118.129:4150",conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("连接nsqd")
	for i := 0; i < 100; i++ {
	err =	prod.Publish("topic",[]byte(strconv.Itoa(i)))
	if err != nil {
		panic(err)
	}
	}
	fmt.Println("生产100个消息")
	 prod.Stop()
	 fmt.Println("生产消息完成 ")

	

}
