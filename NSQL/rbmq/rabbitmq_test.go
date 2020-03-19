package rbmq

//测试单个文件，一定要带上被测试的原文件，如果原文件有其他引用，也需一并带上。
import (
	"sync"
	"testing"
)

var sw sync.WaitGroup

//like  TestXxxxx.....
func TestMq(t *testing.T) {
	sw.Add(3)
	go provider(sw)
	go consumer("queA", "exA", sw)
	go consumer("queB", "exB", sw)
	sw.Wait()
}
