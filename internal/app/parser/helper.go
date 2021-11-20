package parser

import (
	"fmt"
	"math/rand"
	"regexp"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func checkContains(content, pattern string) bool {
	check, err := regexp.MatchString(pattern, content)
	fmt.Println(content)
	if err != nil {
		return false
	}
	return check
}
