func threeSum(nums []int) [][]int {
	triplets := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		pairs := pairSumSortedAllPairs(nums, i+1, -nums[i])
		for _, pair := range pairs {
			triplet := []int{nums[i], pair[0], pair[1]}
			triplets = append(triplets, triplet)
		}
	}

	return triplets
}

func pairSumSortedAllPairs(nums []int, start int, target int) [][]int {
	pairs := [][]int{}
	left, right := start, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			pairs = append(pairs, []int{nums[left], nums[right]})
			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return pairs
}