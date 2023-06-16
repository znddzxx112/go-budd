package utils

// 判断元素是否在数组中
func InStringSlice(needle string, hystack []string) bool {
	for _, item := range hystack {
		if item == needle {
			return true
		}
	}
	return false
}

// 判断元素是否在数组中
func InIntSlice(needle int, hystack []int) bool {
	for _, item := range hystack {
		if item == needle {
			return true
		}
	}
	return false
}

// 删除数组中重复元素
func RemoveRepByLoop(slc []int) []int {
	result := []int{}
	for i := range slc {
		for j := range result {
			if flag := InIntSlice(j, result); flag == false {
				result = append(result, slc[i])
			}
		}
	}
	return result
}
