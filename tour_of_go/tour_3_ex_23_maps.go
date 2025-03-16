// Run in Go browser sandbox to get test data

package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	// fmt.Printf("Words are: %q", words)
	for i := 0; i < len(words); i++ {
		w := words[i]
		count, found := counts[w]
		fmt.Println("count:", count, "found:", found)
		if !(found) {
			counts[w] = 1
		} else {
			counts[w] = counts[w] + 1
		}
	}
	return counts
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

/*
Words are: ["I" "am" "learning" "Go!"]count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
Words are: ["The" "quick" "brown" "fox" "jumped" "over" "the" "lazy" "dog."]count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
Words are: ["I" "ate" "a" "donut." "Then" "I" "ate" "another" "donut."]count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 1 found: true
count: 1 found: true
count: 0 found: false
count: 1 found: true
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
Words are: ["A" "man" "a" "plan" "a" "canal" "panama."]count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 0 found: false
count: 1 found: true
count: 0 found: false
count: 0 found: false
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
*/