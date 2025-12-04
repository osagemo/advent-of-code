package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var AllDirections = []Coordinate{
	{0, 1}, {1, 0}, {1, 1},
	{0, -1}, {-1, 0}, {-1, -1},
	{-1, 1}, {1, -1},
}

func ConcatDigitsToInt(nums ...int) int {
	if len(nums) == 0 {
		panic("no arguments provided")
	}

	var builder strings.Builder
	for _, n := range nums {
		if n < 0 {
			panic(fmt.Sprintf("negative number not allowed: %d", n))
		}
		builder.WriteString(strconv.Itoa(n))
	}

	return MustParseInt(builder.String())
}

func MustParseInts(strings []string) []int {
	ints := []int{}
	for _, s := range strings {
		ints = append(ints, MustParseInt(s))
	}
	return ints
}

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Coordinate struct {
	Row, Col int
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{Row: c.Row + other.Row, Col: c.Col + other.Col}
}

func WithinBounds(matrix [][]rune, coord Coordinate) bool {
	return coord.Row >= 0 && coord.Row < len(matrix) &&
		coord.Col >= 0 && coord.Col < len(matrix[0])
}
