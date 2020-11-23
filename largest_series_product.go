package lsproduct

import (
	"fmt"
)

func LargestSeriesProduct(digits string, n int) (int64, error) {

	if len(digits) < n {

		return -1, fmt.Errorf("span must be smaller than string length")
	}

	if n < 0 {
		return -1, fmt.Errorf("span must be greater than zero")
	}

	l := len(digits)
	end := l - n + 1
	var maxproduct int64 = 0
	if end > l {
		end = l
		maxproduct = 1
	}

	for i, t := range digits[:end] {
		if t < '0' || t > '9' {
			return -1, fmt.Errorf("digits input must only contain digits")
		} else {
			var product int64 = 1
			for _, r := range digits[i : i+n] {
				product *= int64(r - '0')
			}
			if product > maxproduct {
				maxproduct = product
			}
		}
	}
	return maxproduct, nil

}
