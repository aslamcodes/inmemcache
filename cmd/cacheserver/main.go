package main

import (
	"encoding/json"
	"sync"

	"github.com/aslamcodes/inmemcache/cache"
)

func main() {
	c := cache.NewCache[string, []byte]()

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := range 10 {
				val, _ := json.Marshal(j)
				c.Set("i", val)
			}
		}(i)

	}
	wg.Wait()

}
