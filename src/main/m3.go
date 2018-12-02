package main

import (
	"fmt"
)
/**

给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
 */


func lengthOfLongestSubstring(s string) int {

	return len(s)
}

func closure(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main(){
	fmt.Print(lengthOfLongestSubstring("abcabcbb"))
	m := map[int]string{1: "sdklf"}
	m[33] = "OK"
	fmt.Println(m)
	f := closure(10)

	fmt.Println(f(4))

	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("d")
	dd()

}

func dd(){
	var f = [4] func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("Defer fmt: ", i)
		defer func(){
			fmt.Println("Defer fmt: ", i)
		}()
		f[i] = func(){

			fmt.Println("Function seq: ", i)
		}
	}
	for _, n := range f{
		n()
	}
}
/**
0
1
2
3
4
3
4
2
4
1
4
0

 */