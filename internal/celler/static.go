package celler

import (
	"strconv"
	"strings"
)

func Format(value string) string {
	return strings.ToUpper(value)
}

func Convert(i, j int) string {
	if i <= 0 || j <= 0 {
		return ""
	}

	if i > len(alpha) {
		return ""
	}

	row := alpha[i]
	column := strconv.Itoa(j)

	return row + column
}
