package main

func reverse(x int) int {
	// 构造返回值，初始化为0
	ret := 0

	// 采用无限循环遍历一次输入的数据
	for {

		// 如果遍历完，直接跳出循环
		if x == 0 {
			break
		}

		// 取输入值的个位，构造为返回值的个位，然后将返回值增加10倍
		ret = (ret * 10 + x % 10)
		// 将输入值减少10倍
		x = x / 10
	}

	// 如果溢出32位整数，则返回0
	if ret < -2147483648 || ret > 2147483647 {
		return 0
	} else {
		return ret
	}
}

func main() {
	println(reverse(123))
	println(reverse(321))
	println(reverse(-987))
	println(reverse(100))
	println(reverse(-9876))
	println(reverse(1563847412))

	// 输出2^31次方值
	println(2<<30)
}