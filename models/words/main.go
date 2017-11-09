package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	mutil "github.com/cznic/mathutil"
	gopat "github.com/tchap/go-patricia/patricia"
)

const (
	pathToWordFile string = "/var/www/wwf/data/words.txt"
)

type Letters map[byte]int

var (
	wordfile     string
	chars        string
	letters      Letters
	seed         []byte
	permutations []string
)

type Sortable []byte

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	flag.StringVar(&wordfile, "wordfile", pathToWordFile, "Dictionary file to use")
	flag.StringVar(&chars, "chars", "", "Characters to search for")
	flag.Parse()

	if len(chars) > 0 {
		fmt.Println("Using chars:", chars)
		letters = make(map[byte]int, len(chars))
		reg := regexp.MustCompile("[1-4][a-z]")
		found := reg.FindAllString(chars, -1)
		fmt.Println("Found:", found)
		for i := range found {
			letters[byte(found[i][1])] = int(found[i][0])
		}
		fmt.Println("Using letters:", letters)
		seed = makeSlice(letters)
		fmt.Println("Using seed:", seed)
		permutations = makePermutations(seed)
		fmt.Println(permutations)
	}
}

func makeSlice(let Letters) []byte {
	var result []byte
	for c, n := range let {
		for i := 1; i <= n-48; i++ {
			result = append(result, c)
		}
	}
	return result
}

func makePermutations(start Sortable) []string {
	var results []string

	mutil.PermutationFirst(start)
	results = append(results, string(start))
	for mutil.PermutationNext(start) {
		results = append(results, string(start))
	}
	return results
}

func main() {
	createTrieTree(wordfile)
}

func createTrieTree(wf string) {
	trunk := gopat.NewTrie(gopat.MaxPrefixPerNode(20), gopat.MaxChildrenPerSparseNode(26))
	f, err := os.Open(wf)
	if err != nil {
		fatal("Opening word file", err)
	}
	defer f.Close()

	word, _ := regexp.Compile("^[a-zA-Z]+$")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		test := scanner.Text()
		if len(test) < 3 {
			continue
		}
		if word.MatchString(test) {
			t := trunk.Insert(gopat.Prefix(test), true)
			if !t {
				fmt.Println(test)
			}
		}
	}

	/*
		printItem := func(prefix gopat.Prefix, item gopat.Item) error {
			fmt.Printf("%q: %v\n", prefix, item)
			return nil
		}

		trunk.VisitSubtree(gopat.Prefix("calo"), printItem)
	*/
}

func fatal(msg string, err error) {
	fmt.Println(msg)
	fmt.Println(err)
	os.Exit(1)
}
