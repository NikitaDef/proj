package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	copyVal := x
	var reverse int
	for copyVal > 0 {
		splittTail := copyVal % 10
		if reverse == 0 && splittTail == 0 {
			return false
		}
		copyVal = copyVal / 10
		if reverse == 0 {
			reverse = splittTail
		} else {
			reverse = reverse*10 + splittTail
		}
	}
	return x == reverse
}
