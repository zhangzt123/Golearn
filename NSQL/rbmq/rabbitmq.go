/**

sudo docker run -d -p 15672:15672 -p 5672:5672 --name rabbitmq_management  rabbitmq:3.8.3-management

*/

package rbmq

import (
	// "encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var err error
var rec = func() {
	if err := recover(); err != nil {
		log.Fatal(err)

		defer conn.Close()
		defer ch.Close()
	}
}

//

func init() {
	//amqp://<user>:<password>@<ip>:<port>
	conn, err = amqp.Dial("amqp://guest:guest@192.168.118.129:5672")
	if err != nil {
		panic(err)
	}
	fmt.Println("连接rabbitmq成功")
	// defer conn.Close()
	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
	// defer ch.Close()
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
	err = ch.ExchangeBind("exB", "K", "exA", false, nil)
	if err != nil {
		panic(err)
	}
	//创建队列
	queA, err := ch.QueueDeclare("queA", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	queB, err := ch.QueueDeclare("queB", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("绑定前 \n queA:%v \n queB:%v \n", queA, queB)
	//绑定队列
	err = ch.QueueBind(queA.Name, "K", "exA", false, nil)
	if err != nil {
		panic(err)
	}
	err = ch.QueueBind(queB.Name, "K", "exB", false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("绑定后 \n queA:%v \n queB:%v \n", queA, queB)
}

//consumer .
func consumer(que, str string, sw sync.WaitGroup) {
	defer sw.Done()
	defer rec()

	/*
		queue:队列名称。
		consumer:消费者标签，用于区分不同的消费者。
		autoAck:是否自动回复ACK，true为是，回复ACK表示高速服务器我收到消息了。建议为false，手动回复，这样可控性强。
		exclusive:设置是否排他，排他表示当前队列只能给一个消费者使用。
		noLocal:如果为true，表示生产者和消费者不能是同一个connect。
		nowait：是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息。（不推荐使用）
		————————————————
		版权声明：本文为CSDN博主「vrg000」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
		原文链接：https://blog.csdn.net/vrg000/article/details/81165030
	*/
	//fmt.Println(conn.IsClosed())
	a, _ := amqp.Dial("amqp://guest:guest@192.168.118.129:5672")
	c, _ := a.Channel()
	d, err := c.Consume(que, str, true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for x := range d {
		fmt.Println("consumer" + str + ":" + string(x.Body))
		if strings.EqualFold(string(x.Body), "end") {
			fmt.Println("consumer" + str + "接收完成")
			break
		}
	}
	defer conn.Close()
	defer ch.Close()

}

//provider .
func provider(sw sync.WaitGroup) {
	defer sw.Done()
	defer rec()
	//循环20条消息
	for i := 0; i < 20; i++ {
		_ = ch.Tx()
		err = ch.Publish("exA", "K", false, false, amqp.Publishing{Body: []byte(strconv.Itoa(i))})
		if err != nil {
			panic(err)
		}
		_ = ch.TxCommit()
		time.Sleep(1 * time.Second)
		fmt.Println("发送消息:" + strconv.Itoa(i))
	}
	//发消息
	err = ch.Publish("exA", "K", false, false, amqp.Publishing{Body: []byte("end")})
	if err != nil {
		panic(err)
	}
	fmt.Println("发送消息:success")
}
