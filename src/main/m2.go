package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums{
		t := target - v
		tv, e := m[t]
		if e{
			return []int{tv, i}
		}
		m[v] = i
	}

	return nil
}

func main(){
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	m := map[int]int{}
	m[12] = 3
	m[4] = 6
	r, e := m[10]

	fmt.Print(r, e)
}