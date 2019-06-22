package main

import "fmt"

func pattern(pats [][]byte) (ret []string) {
	if len(pats) == 0 {
		return []string{""}
	}

	for _, s := range pattern(pats[1:]) {
		for _, p := range pats[0] {
			ret = append(ret, string(p)+s)
		}
	}
	return ret
}

func letterCombinations(digits string) []string {
	m := map[byte][]byte{'0': []byte{' '}, '1': []byte{}, '2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'}, '4': []byte{'g', 'h', 'i'}, '5': []byte{'j', 'k', 'l'}, '6': []byte{'m', 'n', 'o'}, '7': []byte{'p', 'q', 'r', 's'}, '8': []byte{'t', 'u', 'v'}, '9': []byte{'w', 'x', 'y', 'z'}}

	pats := [][]byte{}
	for i := range digits {
		pats = append(pats, m[digits[i]])
	}

	ret := pattern(pats)
	if ret[0] == "" {
		return []string{}
	}
	return ret
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%+v", letterCombinations(s))
}
