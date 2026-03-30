func twoSum(nums []int, target int) []int {
    pairHash := make(map[int]int)
    for index, num := range nums {
        leftSum := target - num
        if prevIndex, ok := pairHash[leftSum]; ok {
            return []int {prevIndex, index}
        }
        pairHash[num] = index
    }
    return nil
}
