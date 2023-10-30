package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	suffix := make(map[string]*list.List)
	prefix := make(map[string]*list.List)

	for i := 0; i < n; i++ {
		var key string
		fmt.Scanln(&key)
		if suffix[string(key[:3])] == nil {
			suffix[string(key[:3])] = list.New()
		}
		suffix[string(key[:3])].PushBack(key)

		if prefix[string(key[len(key)-3:])] == nil {
			prefix[string(key[len(key)-3:])] = list.New()
		}
		prefix[string(key[len(key)-3:])].PushBack(key)
	}

	for i := 0; i < m; i++ {
		var word string
		fmt.Scanln(&word)

		var pref = string(word[len(word)-3:])
		var suff = string(word[:3])

		if suffix[suff] == nil || prefix[pref] == nil {
			fmt.Println("none")
		} else if suffix[suff].Len() > 1 || prefix[pref].Len() > 1 {
			fmt.Println("ambiguous")
		} else {
			fmt.Println(suffix[suff].Front().Value, prefix[pref].Front().Value)
		}

	}
}
