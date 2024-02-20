package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\t Hello, World\n "
	fmt.Printf("Before Trim Length: %d String:%v\n", len(str), str)
	trim := strings.TrimSpace(str)
	fmt.Printf("After Trim Length: %d String:%v\n", len(trim), trim)
}
