package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func wordSubsets(A []string, B []string) (ret []string) {
	mb := make(map[byte]int, 0)
	for _, b := range B {
		m := map[byte]int{}
		for i := range b {
			m[b[i]]++
		}
		for k, v := range m {
			num, ok := mb[k]
			if !ok {
				mb[k] = v
			} else {
				mb[k] = max(num, v)
			}

		}
	}

	for _, a := range A {
		m := map[byte]int{}
		for i := range a {
			m[a[i]]++
		}
		universal := true
		for k, v := range mb {
			if num, ok := m[k]; !ok || num < v {
				universal = false
			}
		}
		if universal {
			ret = append(ret, a)
		}
	}

	return ret
}

func main() {

}
