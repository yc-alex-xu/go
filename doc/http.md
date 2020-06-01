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

REST是一种架构风格，汲取了WWW的成功经验：无状态，以资源为中心，充分利用HTTP协议和URI协议，提供统一的接口定义，使得它作为一种设计Web服务的方法而变得流行。在某种意义上，通过强调URI和HTTP等早期Internet标准，REST是对大型应用程序服务器时代之前的Web方式的回归。目前Go对于REST的支持还是很简单的，通过实现自定义的路由规则，我们就可以为不同的method实现不同的handle，这样就实现了REST的架构。

note:

HTML标准只能通过链接和表单支持GET和POST。在没有Ajax支持的网页浏览器中不能发出PUT或DELETE命令

Alex; 我的理解是，相比传统的www网站开发，REST　API是与网页分离，就强强调这是一种api，跟本机的函数api唯一区别就是它由http协议承载。

# examples

[code](https://github.com/yc-alex-xu/go/tree/master/src/practise/http)
* get.go: start a http server;start with http://localhost:9999/login?username=John
* form.go: 处理form/MultipartForm/cookie/; start with: http://localhost:9999/

# reference
* [template](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.4.md)
* [JSON](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md)
* [httprouter](https://github.com/julienschmidt/httprouter)