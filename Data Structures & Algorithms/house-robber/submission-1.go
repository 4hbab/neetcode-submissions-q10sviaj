func rob(nums []int) int {
	rob1, rob2 := 0, 0
	for _, currMoney := range nums {
		tmp := max(rob2, rob1 + currMoney)
		rob1 = rob2
		rob2 = tmp
	}
	return rob2
}
