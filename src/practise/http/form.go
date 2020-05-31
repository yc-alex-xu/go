package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func loginpage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.gtpl")
	log.Println(t.Execute(w, nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()
	fmt.Fprintf(w, "email:%v", r.Form["email"])
	fmt.Fprintf(w, "password:%v", r.Form["password"])
}

func main() {
	http.HandleFunc("/", loginpage)  //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	port := os.Getenv("port")
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
