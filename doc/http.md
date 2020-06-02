# prepartion
## install libs
```bash 
$ go get -u github.com/julienschmidt/httprouter
```



## export http port
```bash
$ source export.sh 
```
# REST(REpresentational State Transfer)
什么是RESTful架构：
* 每一个URI代表一种资源；
* 客户端和服务器之间，传递这种资源的某种表现层；
* 客户端通过四个HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"。

note:
* 目前Go对于REST的支持还是很简单的，通过实现自定义的路由规则，我们就可以为不同的method实现不同的handle，这样就实现了REST的架构。
* HTML标准只能通过链接和表单支持GET和POST。在没有Ajax支持的网页浏览器中不能发出PUT或DELETE命令

Alex:相比传统的www网站开发，REST API与网页分离，强调这是一种api，跟本机的函数调用区别就是它由http协议承载。

# examples

[code](https://github.com/yc-alex-xu/go/tree/master/src/practise/http)
* get.go: start a http server;start with http://localhost:9999/login?username=John
* form.go: 处理get/form/MultipartForm/cookie/; start with: http://localhost:9999/
* REST.go  

# reference
* [template](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.4.md)
* [JSON](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md)
* [httprouter](https://github.com/julienschmidt/httprouter)