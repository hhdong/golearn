package main

import (
	"fmt"
	"sort"
)

var (
	drinks = map[string]string{"cocoa": "cocoa奶茶", "b": "bb", "a": "a", "jk": "夹克"}
)

func main() {
	for k, v := range drinks {
		fmt.Printf("key:%s, vlaue:%s\n", k, v)
	}
	keys := make([]string, len(drinks))
	i := 0
	for k, _ := range drinks {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {

		fmt.Printf("key:%s, vlaue:%s\n", k, drinks[k])
	}

}
