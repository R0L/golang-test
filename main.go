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

	// ret := leetcode.MissingNumber([]int{0, 1, 2, 3, 5, 6})
	// fmt.Println(ret)

	// [[7,null],[13,0],[11,4],[10,2],[1,0]]

	var a1, a10, a11, a13, a7 *leetcode.Node

	a7 = &leetcode.Node{
		Val:    7,
		Next:   a13,
		Random: nil,
	}

	a13 = &leetcode.Node{
		Val:    13,
		Next:   a11,
		Random: a7,
	}

	a11 = &leetcode.Node{
		Val:    11,
		Next:   a10,
		Random: a11,
	}

	a10 = &leetcode.Node{
		Val:    10,
		Next:   a1,
		Random: a7,
	}

	a1 = &leetcode.Node{
		Val:    1,
		Next:   nil,
		Random: a7,
	}

	ret := leetcode.CopyRandomList(a7)
	fmt.Println(ret)
}
