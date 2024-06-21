package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {
	var firstLine, secondLine []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 2 && scanner.Scan(); i++ {
		switch i {
		case 1:
			firstLine = numbers(scanner.Text())
		case 2:
			secondLine = numbers(scanner.Text())
		}
	}
	fmt.Println(firstLine)
	fmt.Println(secondLine)
}
