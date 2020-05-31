package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/* 这只是单client 的情况，如果多client 应该用map来表示
 */
var tokenSaved string

/* 起始页
 */
func loginpage(w http.ResponseWriter, r *http.Request) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, token)
	tokenSaved = token
}

/* 处理 file upload 逻辑
 */
func upload(w http.ResponseWriter, r *http.Request) {
	var maxMemory int64 = (32 << 20)
	r.ParseMultipartForm(maxMemory)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("/tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

/*防止多次递交表单,模板中用了隐含字段　token
 */
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()
	token := r.Form.Get("token")
	if token != "" {
		if token != tokenSaved {
			template.HTMLEscape(w, []byte("the form submitted before"))
			return
		}

	} else {
		template.HTMLEscape(w, []byte("wrong form posted"))
		return
	}
	template.HTMLEscape(w, []byte(r.Form.Get("email"))) //输出到客户端
}

func main() {
	http.HandleFunc("/", loginpage) //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	port := os.Getenv("port")
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
