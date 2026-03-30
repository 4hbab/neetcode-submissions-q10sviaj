func hasDuplicate(nums []int) bool {
    tracker := make(map[int]bool)
    for _, num := range nums {
        _, ok := tracker[num]
        if ok {
            return true
        }
        tracker[num] = true
    }
    return false
}
