package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ptopenny/nihon"
)

var DIGRAPH bool = true
var jpmap map[string]string

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		jpmap = nihon.AllHiragana()
	} else if args[0] == "hiragana" {
		jpmap = nihon.RomajiToJapaneseMap
	} else if args[0] == "katakana" {
		jpmap = nihon.KatakanaToJapaneseMap
	}

	var arr []string
	for k, _ := range jpmap {
		if DIGRAPH || (!DIGRAPH && len(k) == 3) {
			arr = append(arr, k)
		}
	}
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(len(arr))
	for _, idx := range p {
		fmt.Printf("%s ?\n", arr[idx])
		for !check(getVal(arr[idx])) {
		}
	}
}

func getVal(val string) string {
	h, _ := jpmap[val]
	return h
}

func check(exp string) bool {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	if text == exp {
		return true
	}
	return false
}
