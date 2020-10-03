package leetcode

/*
XOR 只有某个位置的数字出现奇数次时，该位的掩码才不为 0
此法可以检测出出现一次的位和出现三次的位，为了区分出现一次的数字和出现三次的数字，使用两个位掩码：seen_once 和 seen_twice。

思路是：

仅当 seen_twice 未变时，改变 seen_once。

仅当 seen_once 未变时，改变seen_twice。

作者：LeetCode
链接：https://leetcode-cn.com/problems/single-number-ii/solution/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func singleNumber(nums []int) int {
	first, twice := 0, 0
	for _, v := range nums {
		/*^(XOR),&^ (bit clear AND NOT) 用法可以参考我的例子
		https://github.com/yc-alex-xu/go/blob/master/src/practise/bits/main.go
		*/
		first ^= v
		first &^= twice
		twice ^= v
		twice &^= first
	}
	return first
}
