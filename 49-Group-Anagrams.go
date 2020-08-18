package leetcodego

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		b := bytes(str)
		sort.Sort(b)
		m[string(b)] = append(m[string(b)], str)
	}
	var res [][]string
	for _, arr := range m {
		res = append(res, arr)
	}
	return res
}

type bytes []byte

func (a bytes) Len() int           { return len(a) }
func (a bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bytes) Less(i, j int) bool { return a[i] < a[j] }
