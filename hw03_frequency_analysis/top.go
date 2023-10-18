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
		_, ok := stat[word]

		if ok {
			stat[word]++
		} else {
			stat[word] = 1
		}
	}

	// get all keys from map
	allKeys := make([]string, 0)
	for key := range stat {
		allKeys = append(allKeys, key)
	}

	// sort keys slice by value of stat map
	sort.Slice(allKeys, func(i, j int) bool { return stat[allKeys[i]] > stat[allKeys[j]] })

	// additional sorting on equal counter by keys
	equal := false
	equalStart := 0
	for idx := range allKeys {
		if idx == len(allKeys)-1 && equal {
			sort.Strings(allKeys[equalStart:])
			continue
		}
		if idx == len(allKeys)-1 && !equal {
			continue
		}
		if stat[allKeys[idx]] == stat[allKeys[idx+1]] {
			if !equal {
				equal = true
				equalStart = idx
			}
		} else {
			if equal {
				equal = false
				sort.Strings(allKeys[equalStart : idx+1])
			}
		}
	}

	// return only first 10 if more than it
	if len(allKeys) > 10 {
		return allKeys[:10]
	}

	return allKeys
}
