package captcha

import "strconv"

// type Lookahead func(int, int) int

// InverseCaptcha returns the sum of the list numbers represented by the
// digits based on the add strategy provided
// func InverseCaptcha(input string, lookahead Lookahead) int {
// 	return add(input, lookahead)
// }

// WrapAroundAdd returns the sum of all the digits that match the next digit.
// The list is circular, so the digit after the last digit is the first
// digit in the list.
func WrapAroundAdd(digits string) int {
	return add(digits, 1)
	// if index == size-1 {
	// 	return 0
	// }
	// return index + 1
}

// HalfwayAroundAdd returns the sum of all the digits that match the digit
// halfway around the circular list.
func HalfwayAroundAdd(digits string) int {
	return add(digits, len(digits)/2)
	// return (index + size/2) % size
}

func add(digits string, step int) int {
	sum := 0
	for i := range digits {
		curr := digits[i]
		next := digits[(i+step)%len(digits)]
		if curr == next {
			n, _ := strconv.Atoi(string(curr))
			sum += n
		}
	}

	// for i := 0; i < len(digits); i++ {
	// 	curr := digits[i]
	// 	next := digits[lookahead(len(digits), i)]
	// 	if curr == next {
	// 		n, _ := strconv.Atoi(string(curr))
	// 		sum += n
	// 	}
	// }

	return sum
}
