package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Armstrong(n string) bool {
	power := len(n)
	var sum int64
	for _, digit := range n {
		d, _ := strconv.Atoi(string(digit))
		sum = sum + int64(math.Pow(float64(d), float64(power)))
	}
	if n, _ := strconv.ParseInt(n, 10, 64); n == sum {
		return true
	}
	return false
}

func main() {
	var n string
	for {
		if _, err := fmt.Scanf("%s", &n); err != nil {
			return
		}
		n = strings.TrimSpace(n)
		fmt.Printf("%v\n", Armstrong(n))
	}
}
