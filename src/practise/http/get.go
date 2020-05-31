package main

/*
该程序bind的listen port 由环境变量port指定
$ export port=9999

http://localhost:9999/?url_long=111&url_long=222

可能是package net的限制，如果godoc 也起了一个本地server,启动本程序后，godoc的http server 可能会被杀死

*/
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello from %v !", os.Getenv("USER")) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	port := os.Getenv("port")
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
