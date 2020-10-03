package main

import (
	"fmt"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	dfs(s, 0, 0, [4]int{}, &ans)
	return ans
}

/*
ip 存的就是中间结果
*/
func dfs(s string, pos, idx int, ip [4]int, ans *[]string) {
	//已经遍历完了字符串
	if idx == 4 && pos == len(s) {
		t := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
		*ans = append(*ans, t)
		return
	}
	//本次试探失败
	if idx == 4 || pos == len(s) {
		return
	}
	//不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[pos] == '0' {
		ip[idx] = 0
		dfs(s, pos+1, idx+1, ip, ans)
		return
	}
	t := 0
	for i := pos; i < len(s); i++ {
		t = t*10 + int(s[i]-'0')
		if t > 0 && t <= 0xFF {
			ip[idx] = t
			dfs(s, i+1, idx+1, ip, ans)
		} else {
			//提前终结
			break
		}
	}
	return
}
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("1111"))
	fmt.Println(restoreIpAddresses("010010"))
	fmt.Println(restoreIpAddresses("101023"))
}
