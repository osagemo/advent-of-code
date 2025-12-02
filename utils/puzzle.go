package utils

import (
	"fmt"
	"strings"
	"time"
)

var Debug = false

func PrepareInput(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Trim(input, "\n")
}

func RunTimed(name string, fn func() int) {
	start := time.Now()
	fmt.Printf("%s: %d\n", name, fn())
	fmt.Println(time.Since(start))
}

func DebugLog(format string, a ...any) {
	if Debug {
		fmt.Printf(format, a...)
	}
}
