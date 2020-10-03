leetCode

# 代码原则
逻辑自然，易懂

# unit test

unit test 一般放在main_test.go中，如果比较太复杂就在 func main() 中执行但不做比较．

* vs code　可以选择test casess直接执行．
* 如果是命令行
```bash
go/src/leetcode$ go test -timeout 30s leetcode/010
ok  	leetcode/010	0.002s
```