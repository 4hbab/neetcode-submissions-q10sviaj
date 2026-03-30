func maxArea(heights []int) int {
	max_area := 0
	left, right := 0, len(heights)-1
	for left < right {
		currMax := min(heights[right], heights[left])*(right-left)
		max_area = max(max_area, currMax)
		if heights[right] < heights[left] {
			right--
		} else if heights[right] > heights[left] {
			left++
		} else {
			left++
			right--
		}
	}
	return max_area
}
