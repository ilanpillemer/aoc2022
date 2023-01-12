package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("...day 25...")
	b, _ := os.ReadFile("INPUT")
	lines := strings.Split(string(b), "\n")
	var total int64
	fmt.Println(lines)
	//	fmt.Println("2=-01 ==> ", fromSnafu("2=-01"))
	//	fmt.Println("1121-1110-1=0 ==> ", fromSnafu("1121-1110-1=0"))
	fmt.Println("2==0=0===02--210---1 ==> ", fromSnafu("2==0=0===02--210---1"))
	//os.Exit(0)
	for _, line := range lines {
		total = total + fromSnafu(line)
	}
	fmt.Println(total)

	fmt.Println(toSnafu(total))
}

func fromSnafu(num string) int64 {
	var total int64

	for i, j := len(num)-1, 0.0; i >= 0; i, j = i-1, j+1 {
		col := int64(math.Pow(5, j))
		//	fmt.Printf("%d * %d\n", f(num[i]), col)
		total = total + (f(num[i]) * col)
	}
	return total
}

func toSnafu(num int64) string {
	base5 := strconv.FormatInt(num, 5)
	fmt.Println(base5)
	carry := 0
	_ = carry
	str := ""
	_ = str
	for i, j := len(base5)-1, 0.0; i >= 0; i, j = i-1, j+1 {
		n := atoi(string(base5[i]))
		n = n + carry
		carry = 0
		switch n {
		case 0:
			str = "0" + str
		case 1:
			str = "1" + str
		case 2:
			str = "2" + str
		case 3:
			str = "=" + str
			carry = 1
		case 4:
			str = "-" + str
			carry = 1
		case 5:
			str = "0" + str
			carry = 1
		}
		//col := int64(math.Pow(5, j))
		//fmt.Printf("%d %d\n", n, col)
	}

	return str
}

func f(x byte) int64 {
	m := map[byte]int64{
		'2': 2,
		'1': 1,
		'-': -1,
		'=': -2,
	}
	return m[x]
}

func atoi(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}
