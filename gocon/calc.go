package gocon

// FindTheBiggest 使用GoConvey框架进行测试
func FindTheBiggest(sourceList []int) int {
	if len(sourceList) == 0 {
		return -1
	}
	target := sourceList[0]
	for _, val := range sourceList {
		if val > target {
			target = val
		}
	}

	return target
}

// FindIfIn find if the target number in the array
func FindIfIn(sourceList []int, target int) bool  {
	for _, val := range sourceList {
		if target == val {
			return true
		}
	}

	return false
}
