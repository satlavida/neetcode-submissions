func topKFrequent(nums []int, k int) []int {
	if(len(nums) < k) { 
		return []int{}
	}

	freq := make(map[int]int, len(nums))

	for _, v := range nums {
		freq[v]++
	}

	buckets := make([][]int, len(nums)+1)

	for n,count := range freq {
		buckets[count] = append(buckets[count], n)
	}

	result := make([]int, 0, k)

	for c := len(buckets)-1; c >= 0 && len(result) < k; c-- {
		result = append(result, buckets[c]...)
	}
	return result
}
