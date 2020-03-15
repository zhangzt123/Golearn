package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

const form = `
	<html><body>
		<form action="http://localhost:8080/culhtml" method="post" name="bar">
			<p>数字x</p><input type="text" name="numx" value="${name}" /></br>
			<p>数字y</p><input type="text" name="numy" value="${name}" /></br>
			<p>结果</p><input type="text" name="numy" value="${ren}" /></br>
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

// Myhandle .
// type Myhandle struct {
// }

// func (Myhandle *Myhandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {

// }

// cul  x+y .
func cul(x, y int) (ret int, Err string) {

	ret = x + y
	return
}

// culhtml .
func culhtml(w http.ResponseWriter, req *http.Request) {
	numx, err := strconv.Atoi(req.PostFormValue("numx"))
	if err != nil {
		panic(err)
	}
	numy, Err := strconv.Atoi(req.PostFormValue("numy"))
	if Err != nil {
		panic(Err)
	}
	x, _ := cul(numx, numy)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	str := strings.Replace(form, "${name}", req.PostForm.Get("numx"), 1)
	str = strings.Replace(str, "${name}", req.PostForm.Get("numy"), 1)
	str = strings.Replace(str, "${ren}", strconv.Itoa(x), 1)
	w.Write([]byte(str))
}

// home .
func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//str := req.URL.Query().Get("name")
	//relpace的最后表示替换次数
	str := strings.Replace(form, "${name}", "", 2)
	str = strings.Replace(str, "${ren}", "", 1)
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/culhtml", logPanics(culhtml))
	http.HandleFunc("/home", logPanics(home))
	//var m Myhandle
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//HandleFnc .
type HandleFnc func(http.ResponseWriter, *http.Request)

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}
