package main

import (
	"fmt"
	"time"

	"github.com/aslamcodes/inmemcache/cache"
)

func main() {
	c := cache.NewCache[string, string]()

	c.SetWithTTL("session", "abcdefg", 2*time.Second)

	if val, ok := c.Get("session"); ok {
		fmt.Println("immediately:", val, ok)
	} else {
		fmt.Println("immediately: miss (unexpected!)")
	}

	time.Sleep(3 * time.Second)

	// read after expiry — should be a miss
	if val, ok := c.Get("session"); ok {
		fmt.Println("after ttl:", val, ok, "(unexpected!)")
	} else {
		fmt.Println("after ttl: miss (expected)")
	}

}
