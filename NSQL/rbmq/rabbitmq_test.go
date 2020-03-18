package rbmq
//测试单个文件，一定要带上被测试的原文件，如果原文件有其他引用，也需一并带上。
import (
	"fmt"
	"testing"
)

func Testmq(t *testing.T) {
	Consumer()
	fmt.Println("hahahhaah ")
}
