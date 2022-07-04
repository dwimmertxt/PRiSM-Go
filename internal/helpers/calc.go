package helpers

func Centre(a int, b int) int {
	centre := (a / 2) - (b / 2)
	return centre
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Modulo(a, b int) int {
	return ((a % b) + b) % b
}

func NumberToBase(n, b int) int {
	if n == 0 {
		return 0
	}
	var digits []int
	var condition bool = true
	for condition {
		digits = append(digits, Modulo(n, b))
		n = n / b
		if n == 0 {
			condition = false
		}
	}
	return Sum(digits)
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}