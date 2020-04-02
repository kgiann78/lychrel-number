package lychrel_number

import (
	"fmt"
	"math/big"
	"strings"
)

func converges(n int) int {
	iteration := 0
	temp := big.NewInt(int64(n))

	for i := 0; i <= 1000 && !isPalindrome(temp); i++ {
		iteration += 1
		num, ok := reverse(temp)
		if !ok {
			fmt.Println("Error? ", ok)
			break
		} else {
			temp.Add(temp, num)
		}
	}

	return iteration
}

func reverse(n *big.Int) (*big.Int, bool) {
	digits := n.Text(10)
	pDigits := make([]string, len(digits))
	lastIndex := len(digits) - 1

	for i := 0; i < len(digits); i++ {
		pDigits[i] = string(digits[lastIndex-i])
	}

	str := strings.TrimLeft(strings.Join(pDigits, ""), "0")
	num, ok := new(big.Int).SetString(str, 0)

	return num, ok
}

func isPalindrome(n *big.Int) bool {
	digits := n.Text(10)
	lastIndex := len(digits) - 1

	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[lastIndex-i] {
			return false
		}
	}

	return true
}
