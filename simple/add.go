package simple

import "fmt"

func Add(a, b int) int {
	if a == 1 {
		fmt.Println("a is 1")
	}
	if b == 2 {
		fmt.Println("b is 2")
	}
	return a + b
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1len := len(nums1)
	nums2len := len(nums2)
	stack := []int{}
	result := []int{}
	top := -1
	indexMap := make(map[int]int)
	for i := nums2len-1; i >= 0; i-- {
		indexMap[nums2[i]] = i
		for top >= 0 {
			if stack[top] <= nums2[i] {
				top -= 1
			} else {
				break
			}
		}
		fmt.Println(stack)
		fmt.Println(top)
		if top >= 0 {
			result = append([]int{stack[top]}, result...)
		} else {
			result = append([]int{-1}, result...)
		}
		top += 1
		if len(stack) - 1 >= top {
			stack[top] = nums2[i]
		} else {
			stack = append(stack, nums2[i])
		}
	}

	fmt.Println(result)

	for i := 0; i < nums1len; i++ {
		nums1[i] = result[indexMap[nums1[i]]]
	}

	return nums1
}

func rangeAndLen(s string) int {
	for _, v := range s {
		fmt.Println(v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	return 0
}
