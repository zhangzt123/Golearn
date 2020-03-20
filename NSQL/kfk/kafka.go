package kfk

import (
	//连接kafka使用
	//Sarama is a Go library for Apache Kafka 0.8, and up
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	//Cluster extensions for Sarama, the Go client library for Apache Kafka 0.9
	//"github.com/bsm/sarama-cluster"
)

var conf *sarama.Config
var client sarama.Client
var err error
var rec = func() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}
var kfkaddr string = "192.168.118.129:9092"

func init() {
	conf = sarama.NewConfig()
	/*
		type RequiredAcks int16
		NoResponse RequiredAcks = 0  不等待broker确认返回消息
		WaitForLocal RequiredAcks = 1 leader确认消息后其他follower 异步接受消息
		WaitForAll RequiredAcks = -1 leader和其他follower都确认接受消息
	*/
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Return.Successes = true
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Consumer.Return.Errors = true

	client, err = sarama.NewClient([]string{kfkaddr}, conf)
	if err != nil {
		panic("连接client失败panic：" + err.Error())
	}
	fmt.Println("连接成功！！！")
	broks := client.Brokers()
	for _, brok := range broks {
		fmt.Println(brok.Addr())
	}

}

//
func consumer(sw sync.WaitGroup) {
	defer sw.Done()
	defer rec()
	consume, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic("连接consume失败panic：" + err.Error())
	}
	sclieconsumes, err := consume.Partitions("Mytopic")
	if err != nil {
		panic("获取consume partition IDs 失败panic：" + err.Error())
	}
	for _, sclieconsume := range sclieconsumes {
		partitionconsume, err := consume.ConsumePartition("Mytopic", sclieconsume, sarama.OffsetNewest)
		if err != nil {
			panic("根据consume partition ID 获取channel 失败panic：" + err.Error())
		}
		consunemsg := <-partitionconsume.Messages()
		fmt.Printf("partition ID:%v ; msg:%v \n ", consunemsg.Partition, string(consunemsg.Value))
	}

}

func provider(sw sync.WaitGroup) {
	defer sw.Done()
	defer rec()
	// AsyncProducer, err := sarama.NewAsyncProducer([]string{kfkaddr}, conf)
	AsyncProducer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic("连接AsyncProducer失败panic：" + err.Error())
	}
	messageer := AsyncProducer.Input()
	msg := sarama.ProducerMessage{}
	msg.Topic = "Mytopic"
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.StringEncoder("this is a message .")
	messageer <- &msg
	AsyncProducer.AsyncClose()
	_ = <-AsyncProducer.Successes()
	fmt.Println("发送success !!!")
}
