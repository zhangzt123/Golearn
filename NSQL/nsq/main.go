// nsq go语言开发的消息对列
//具体安装查看官方文档 此处使用docker镜像
//https: //nsq.io/overview/quick_start.html
/**
docker pull nsqio/nsq:v1.2.0
docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

sudo docker run -d --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq:v1.2.0 /nsqd \
    --broadcast-address=192.168.118.129 \
	--lookupd-tcp-address=192.168.118.129:4160
	
*/

package main

import 
	(
		nsq "github.com/nsqio/go-nsq"
	)

func main() {
	//生产者消费者 topic 
}