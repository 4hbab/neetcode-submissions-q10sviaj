func twoSum(nums []int, target int) []int {
    pairHash := make(map[int]int)
    for index, num := range nums {
        leftSum := target - num
        if _, ok := pairHash[leftSum]; ok {
            return []int {pairHash[leftSum], index}
        }
        pairHash[num] = index
    }
    return nil
}
