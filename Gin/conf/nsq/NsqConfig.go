package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/zhangzt123/Golearn/Gin/conf"
)

var nsqconf *nsq.Config
var addr string

func init() {
	nsqconf = nsq.NewConfig()
	//nsqconf
	addr = conf.Conf.GetString("nsq.addr")
}

func NewNsqConnInstence() *nsq.Conn {
	conn := nsq.NewConn(addr, nsqconf, nil)
	return conn
}

//func NewNsqConsumeInstence() *nsq.Consumer{
//
//}
//
//
//func NewNsqProvideInstence() *nsq.Producer{
//
//}
