Go 提供了一个低级的分析 API runtime/pprof ，但如果你在开发一个长期运行的服务，使用更高级的 net/http/pprof 包会更加便利。你只需要在代码中加入 import _ "net/http/pprof" ，它就会自动注册所需的 HTTP 处理器（Handler） 。
