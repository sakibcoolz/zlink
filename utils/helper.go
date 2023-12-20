package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/starryalley/go-julianday"
)

func ToCharStr(i int) string {
	return fmt.Sprintf("%c", ('A' - 1 + i))
}

func UrlPath() string {
	j := julianday.Date(time.Now())

	return ToEncode(int(j))
}

func ToEncode(val int) string {
	ints := splitToGigits(val)

	strUrl := IntToStringEncode(ints)

	miliUrl := IntToStringEncode(splitToGigits(MiliSeconds()))

	strUrl = append(strUrl, miliUrl...)

	return strings.Join(strUrl, "")
}

func IntToStringEncode(ints []int) []string {
	back, forword := 0, 0
	data := make([]int, 0)

	strslice := make([]string, 0)
	for i := 0; i < len(ints); i = i + 2 {
		if len(ints) <= i+1 {
			forword = 26
		} else {
			forword = ints[i+1]
			if forword == 0 {
				forword = 26
			}
		}
		str := fmt.Sprintf("%d%d", ints[i], forword)
		data = append(data, ints[i], forword)
		back, _ = strconv.Atoi(str)

		if 26 >= back {
			strslice = append(strslice, ToCharStr(back))
		} else {
			strslice = append(strslice, ToCharStr(SwapAndSub(data[0], data[1])))
		}
		data = make([]int, 0)
	}

	return strslice
}

func SwapAndSub(i, j int) int {
	if i < j {
		return j - i
	}

	return i - j
}

func MiliSeconds() int {
	return int(time.Now().UnixNano() / int64(time.Millisecond))
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToGigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)

	return ret
}
