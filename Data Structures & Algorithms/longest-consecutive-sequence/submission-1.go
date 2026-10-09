func longestConsecutive(nums []int) int {
    set := make(map[int]bool, len(nums))
    for _, v := range nums {
        set[v] = true
    }

    longest := 0
    for v := range set {
        if set[v-1] {
            continue
        }
        length := 1
        for set[v+length] {
            length++
        }
        longest = max(longest, length)
    }
    return longest
}