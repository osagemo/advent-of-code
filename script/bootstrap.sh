#!/bin/bash

if [ -z "$1" ] || ! [[ $1 =~ ^[0-9]+$ ]]; then
    echo "Usage: $0 <day-number>"
    exit 1
fi

day_without_zero=$1
day=$(printf "%02d" "$1")

# Create folder
mkdir "${day}"

# Create dayXX.go
cat <<EOL >"${day}"/day"${day}".go
package main

import (
	_ "embed"
	"fmt"

	"github.com/osagemo/advent-of-code/utils"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
    input := utils.PrepareInput(input)
    fmt.Println("Day ${day_without_zero}")
    utils.RunTimed("Part 1", func() int { return Part1(input) })
    utils.RunTimed("Part 2", func() int { return Part2(input) })
}
EOL

# Create dayXX_test.go
cat <<EOL >"${day}"/day"${day}"_test.go
package main

import (
	"fmt"
	"testing"
)

const input1 = \`\`

func TestDay${day_without_zero}Part1(t *testing.T) {
	result := Part1(input1)
	expected := -1

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay${day_without_zero}Part2(t *testing.T) {
	result := Part2(input1)
	expected := -1

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
EOL

# Create an empty input.txt
touch "${day}"/input.txt
