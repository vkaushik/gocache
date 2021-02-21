package main

import (
	"fmt"

	"github.com/vkaushik/gocache"
	"github.com/vkaushik/gocache/lru"
)

func main() {
	var numbersCache gocache.Cache = lru.NewCache(3)
	numbersCache.Set("zero", 0)
	numbersCache.Set("one", 1)
	numbersCache.Set("two", 2)
	numbersCache.Set("three", 3) // zero removed

	if value, cacheFound := numbersCache.Get("zero"); cacheFound {
		fmt.Println("error: { key:zero } is found with value:", value)
	} else {
		fmt.Println("success: { key:zero } is not found.")
	}

	if value, cacheFound := numbersCache.Get("one"); cacheFound {
		fmt.Println("success: { key:one } is found with value:", value)
	} else {
		fmt.Println("error: { key:one } is not found.")
	}

	numbersCache.Set("four", 4) // two removed

	if value, cacheFound := numbersCache.Get("two"); cacheFound {
		fmt.Println("error: { key:two } is found with value:", value)
	} else {
		fmt.Println("success: { key:two } is not found.")
	}

	fmt.Println()
}
