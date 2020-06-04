# build　image
It a little different from [Dev env](docker.md), since it needn't recompilation when code change. 

一般的Dockerfile都会　go get 所有的库，在images中编译，这对CI来说有一定必要，但会增加image体积.通常的解决方案是分多个stage，
```Dockerfile
FROM golang
WORKDIR /src
COPY hello.go .
RUN go build hello.go
FROM ubuntu　　＃开始stage 1
COPY --from=0 /src/hello .  #copy stage ０生成的文件到本stage
CMD ["./hello"]
```
不过最简单的方法还是本地静态编译生成可执行文件

```Dockerfile
FROM scratch
LABEL maintainer="xuyc@sina.com"
ADD . /
EXPOSE 80
CMD ["/myapp"]
```
只要开发的环境与运行的环境ABI兼容，这样直接copy二进制文件肯定是没问题的。

```bash
$ script/create_image.sh 
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
justware/myapp      1.0                 101691d27dfc        8 minutes ago       11.7MB
```
# run image
```bash
$ docker run -it --rm  -p 80:80  justware/myapp:1.0
2020/06/03 13:27:22.666 [I] [asm_amd64.s:1373]  http server Running on http://:80
2020/06/03 13:27:30.443 [D] [server.go:2807]  |     172.17.0.1| 200 |   7.677614ms|   match| GET      /     r:/
2020/06/03 13:27:30.498 [D] [server.go:2807]  |     172.17.0.1| 304 |    221.102µs|   match| GET      /static/js/reload.min.js
```
# publish
```bash
$ script/docker_publish.sh 
The push refers to repository [docker.io/justware/myapp]
1bcc18978517: Pushed 
1.0: digest: sha256:af2e736d15bf832279d5c560c449fb252988586a534c88203c20da69d2009e44 size: 528
```
# run in k8s
tested too, since the node's OS is Linux too.

# CI/CD
* [github action](https://help.github.com/cn/actions/language-and-framework-guides/publishing-docker-images)
* [用 GitHub Action 构建一套 CI/CD 系统](https://cloud.tencent.com/developer/article/1624786)

