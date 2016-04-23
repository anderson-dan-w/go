package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readIn(fname string) string {
	text, err := ioutil.ReadFile(fname)
	check(err)
	return string(text) // []byte otherwise
}

func getWordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, substr := range strings.Fields(s) {
		// defaultdict is the... default!
		count[substr] += 1
	}
	return count
}

func getFollow(s string) map[string][]string {
	prev := ""
	follow := make(map[string][]string)
	for _, substr := range strings.Fields(s) {
		follow[prev] = append(follow[prev], substr)
		prev = substr
	}
	// prevent terminations: last word cycles back to initial word
	follow[prev] = append(follow[prev], "")

	return follow
}

func printRand(follow map[string][]string) {
	word := "" // I know I put this in map in getFollow
	for i := 0; i < 75; i++ {
		fmt.Printf("%v ", word)
		r := rand.Intn(len(follow[word]))
		word = follow[word][r]
	}
}

func main() {
	s := "I am learning go, I think"
	fmt.Println("WORD COUNTS")
	for k, v := range getWordCount(s) {
		fmt.Printf("%v : %v\n", k, v)
	}

	fmt.Println("\nRandom bigrams")
	alice := readIn("data/alice.txt")

	// if we want new random every time...
	rand.Seed(time.Now().Unix())

	printRand(getFollow(alice))
	fmt.Println()

}
