/*
 * @Author: zzt
 * @Date: 2020-03-05 11:34:31
 * @LastEditTime: 2020-03-05 15:25:32
 * @LastEditors: Please set LastEditors
 * @Description: 读取文件并写入到write.txt中
 * @FilePath: \GoProject\src\github.com\zhangzt123\Golearn\learn1\learn1.go
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// "bufio"
// "io"

// "os"

/**
 * @description:
 * @param {type}
 * @return:
 */
func main() {

	// file, err := os.OpenFile("./learn1.go", os.O_RDONLY, 777)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// ioreader := bufio.NewReader(file)
	// // bt := make([]byte, 1024)
	// for {
	// 	bt, err := ioreader.ReadBytes('\n')
	// 	fmt.Printf("%v", string(bt))
	// 	if err == io.EOF {
	// 		os.Exit(0)
	// 	}
	// 	if err != nil {
	// 		os.Exit(0)
	// 	}
	// }

	// fmt.Println("--------------------------------------------------------")

	// bt1, _ := ioutil.ReadFile("./learn1.go")
	// fmt.Printf("%v", string(bt1))

	file, err := os.Open("./learn1.go")
	file1, _ := os.OpenFile("./write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("发生错误:%v", err)
		return
	}
	defer file.Close()
	defer file1.Close()
	bufreader := bufio.NewReader(file)
	for {
		bt, err := bufreader.ReadBytes('\n')
		file1.Write(bt)

		if err == io.EOF {
			fmt.Printf("读取完成:%v", err)
			return
		}
		if err != nil {
			fmt.Printf("发生错误:%v", err)
			return
		}

	}

}

/*
Your workspace is misconfigured: go [-e -json -compiled=true -test=true -export=false -deps=true -find=false -- ./]: exit status 1: go: cannot find main module; see 'go help modules'
. Please see https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md for more information or file an issue (https://github.com/golang/go/issues/new) if you believe this is a mistake.
*/
