package main

import (
	"fmt"
	"github.com/R0L/golang-test/leetcode"
)

func main() {
	//piles := []int{30,11,23,4,20}
	//H := 6
	//ret := leetcode.MinEatingSpeed(piles, H)
	//fmt.Printf("%v", ret)

	//ret := leetcode.MySqrt(83)
	//fmt.Println(ret)

	//ret := leetcode.BulbSwitch(3)
	//fmt.Println(ret)

	//ret := leetcode.CanTransform("RXXLRXRXL", "XRLXXRRLX")
	//fmt.Println(ret)

	//ret := leetcode.IsAnagram("qerr", "reqr")
	//fmt.Println(ret)

	//ret := leetcode.GroupAnagrams([]string{"", ""})
	//ret := leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	//fmt.Println(ret)

	//ret := leetcode.MajorityElement([]int{6, 5, 5})
	//fmt.Println(ret)

	//ret := leetcode.SingleNumber([]int{1, 1, 2, 3, 4, 3, 4})
	//fmt.Println(ret)

	ret := leetcode.MissingNumber([]int{0, 1, 2, 3, 5, 6})
	fmt.Println(ret)
}
