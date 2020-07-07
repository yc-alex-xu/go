# sort

![sort](images/sort.png)

# backtrace
e.g.子集问题
```go
func subsets(nums []int)[][]int {
    if len(nums) == 0{
        return [][]int{}
    }
    temp = [][]int{}
    return backTrace(0, nums, temp);
    
}

func backTrace(start int, nums []int,temp [][]int) (result [][]int){
    base case处理
    //选择过程
    for 循环选择{
        选择
        backTrace(递归);
        撤销选择
    }
}
```    