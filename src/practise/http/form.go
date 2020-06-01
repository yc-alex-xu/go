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

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

/* 这只是单client 的情况，如果多client token应该用map来表示,也许应该放到redis中
 */
var tokenSaved string

/* 起始页
 */
func homepage(w http.ResponseWriter, r *http.Request) {
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
	if cookie, _ := r.Cookie("email"); cookie == nil {
		http.Redirect(w, r, "/", 302)
		return
	} else {
		fmt.Fprint(w, cookie)
	}

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
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	e := r.Form.Get("email")
	cookie := http.Cookie{Name: "email", Value: e, Expires: expiration}
	http.SetCookie(w, &cookie)
	template.HTMLEscape(w, []byte(e)) //输出到客户端
}

func main() {
	http.HandleFunc("/", homepage) //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	port := os.Getenv("port")
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
