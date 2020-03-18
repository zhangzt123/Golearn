/**

sudo docker run -d -p 15672:15672 -p 5672:5672 --name rabbitmq_management  rabbitmq:3.8.3-management

*/

package rbmq

import (
	// "encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var err error
var rec = func() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}

//

// func init() {

// }

//Consumer .
func Consumer() {
	//defer rec()
	//amqp://<user>:<password>@<ip>:<port>
	conn, err := amqp.Dial("amqp://guestguest@192.168.118.129:5672")
	if err != nil {
		panic(err)
	}
	fmt.Println("连接rabbitmq成功")
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	//申请交换机
	//绑定交换机使交换机间进行source 转发 destination ；
	err = ch.ExchangeDeclare("exA", "topic", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare("exB", "topic", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//source  交换机绑定到destination 交换机上 通过key匹配
	err = ch.ExchangeBind("exA", "K", "exA", false, nil)
	if err != nil {
		panic(err)
	}
	//创建队列
	queA, err := ch.QueueDeclare("myqueue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	queB, err := ch.QueueDeclare("myqueue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("queA:%v \n queB:%v \n", queA, queB)
	//绑定队列
	err = ch.QueueBind(queA.Name, "K", "exA", false, nil)
	if err != nil {
		panic(err)
	}
	err = ch.QueueBind(queB.Name, "K", "exB", false, nil)
	if err != nil {
		panic(err)
	}

	//发不消息
	_ = ch.Tx()
	err = ch.Publish("exA", "K", false, false, amqp.Publishing{Body: []byte("helloworld")})
	if err != nil {
		panic(err)
	}
	_ = ch.TxCommit()
	time.Sleep(100 * time.Second)

}

func provider() {
	fmt.Println("")
}
