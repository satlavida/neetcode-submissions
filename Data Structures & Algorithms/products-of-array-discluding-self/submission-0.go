func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	// 1 2 4 6
	// 1 1 2 8

	suffix := 1
	for i := n-1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i] 
	}
	//suffix = 1
	//       8
	
	//suffix = 6
	//        12 8

	//suffix = 6 * 4 = 24
	//     24 12 8

	//suffix = 24 * 2 = 48
	// 48 24 12 8
	return result
}
