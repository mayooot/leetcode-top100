package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	m := make(map[string][]string, len(strs))
	for _, s := range strs {
		ss := sortString(s)
		m[ss] = append(m[ss], s)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

func main() {
	fmt.Println(sortString("bac"))
}
