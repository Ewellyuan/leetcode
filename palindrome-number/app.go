package main

import("fmt")

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var ret bool = true
	i := [...]int{}
	a := i[:]

	// 采用无限循环遍历一次输入的数据
	for {
		// 如果遍历完，直接跳出循环
		if x == 0 {
			break
		}

		// 构造每位数的数组
		a = append(a, x%10)
		// 将输入值减少10倍
		x = x / 10
	}
	
	// 遍历每位数数组
	for b:=0; b<(len(a)/2); b++ {
		if a[b] != a[len(a)-1-b] {
			ret = false
		}
	}

	return ret
}

func main() {
	fmt.Println(100,isPalindrome(100))
	fmt.Println(111,isPalindrome(111))
	fmt.Println(-111,isPalindrome(-111))
	fmt.Println(22,isPalindrome(22))
	fmt.Println(121,isPalindrome(121))
	fmt.Println(1221,isPalindrome(1221))
	fmt.Println(1234,isPalindrome(1234))
	fmt.Println(123454321,isPalindrome(123454321))
}