package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	reSymbols = regexp.MustCompile(`(\t|\n|\. |, |!|"|;| - )`)
	reSpace   = regexp.MustCompile(`\s+`)
)

func Top10(s string) []string {
	var stat map[string]int

	// skip empty string
	if s == "" {
		return nil
	}

	// clean string from unnecessary symbols and split
	cleanedS := reSpace.ReplaceAllString(reSymbols.ReplaceAllString(strings.ToLower(s), " "), " ")
	words := strings.Fields(cleanedS)

	// init map which will be used as counter
	stat = make(map[string]int)

	// iterate on words slice and increase counter
	for _, word := range words {
		stat[word]++
	}

	// get all keys from map
	allKeys := make([]string, 0)
	for key := range stat {
		allKeys = append(allKeys, key)
	}

	// sort keys slice by value of stat map
	sort.Slice(allKeys, func(i, j int) bool { 
	a, b := allKeys[i], allKeys[j]
	if stat[a] == stat[b] {
			return a < b
		}
	return stat[a] > stat[b]
	 })

	// return only first 10 if more than it
	if len(allKeys) > 10 {
		return allKeys[:10]
	}

	return allKeys
}
