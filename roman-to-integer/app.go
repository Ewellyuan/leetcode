package main

import("fmt")

func romanToInt(s string) int {
	ret := 0
	pre := ""

	for i, ch := range s {

		strTmp := string(ch)

		if pre != "" {
			strTmp = pre + strTmp
			pre = ""
		}

		switch strTmp {
			case "C":
			case "X":
			case "I":
		}

		switch strTmp {
			case "M":
				ret += 1000
			case "CM":
				ret += 900
			case "D":
				ret += 500
			case "CD":
				ret += 400
			case "C":
				ret += 100
			case "XC":
				ret += 90
			case "L":
				ret += 50
			case "XL":
				ret += 40
			case "X":
				ret += 10
			case "IX":
				ret += 9
			case "V":
				ret += 5
			case "IV":
				ret += 4
			case "I":
				ret += 1
			default:
				ret += 0
		}
	}

	return ret
}

func doPrint(s string) {
	fmt.Println(s, romanToInt(s))
}

func main()  {
	// romanToInt("VIII")
	doPrint("III")
	doPrint("IV")
	// doPrint("VII")
	// doPrint("IX")
	// doPrint("LVIII")
	// doPrint("MCMXCIV")
}