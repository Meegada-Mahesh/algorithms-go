package main

import (
	"fmt"
)

func naiveMatching(text string, pattern string) {
	fmt.Println("Text and Pattern Values are ", text, pattern)
	var j int
	m := len(text) - len(pattern)
	for i := 0; i <= m; i++ {
		j = 0

		for j < len(pattern) {
			if text[i+j] != pattern[j] {
				break
			} else {
				j++
				if j == len(pattern) {
					fmt.Println("Pattern Matching Found at ", i)
				}
			}
		}
	}
}

func main() {
	txt := "MAHESHisaHumanandHeisAWESOme"
	patt := "ahe"

	naiveMatching(txt, patt)
}
