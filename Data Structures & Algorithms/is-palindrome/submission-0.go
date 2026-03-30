func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isASCIIAlphaNumeric(rune(s[left])) {
			left++
			continue
		}
		if !isASCIIAlphaNumeric(rune(s[right])) {
			right--
			continue
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
func isASCIIAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
