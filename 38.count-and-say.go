package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 5
	output := countAndSay(count)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start

// 1
// 11
// 21
// 1211
// 111221
var m map[int]string = map[int]string{
	1: "1",
	2: "11",
}

func countAndSay(n int) string {
	s := countAndSayImpl(n)
	return s
}

func countAndSayImpl(n int) string {
	if v, ok := m[n]; ok {
		return v
	}

	if n <= 1 {
		return "1"
	}

	s := countAndSayImpl(n - 1)

	r := ""
	l := 0

	var p int
	for p = 1; p < len(s); p++ {
		if s[l] != s[p] {
			count := p - l
			r = r + strconv.Itoa(count) + string(s[l])
			l = p
		}
	}

	if l != p {
		count := p - l
		r = r + strconv.Itoa(count) + string(s[l])
	}

	m[n] = r

	return r
}

// @lc code=end
