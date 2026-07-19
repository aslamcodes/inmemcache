package cache_test

import (
	"testing"

	"github.com/aslamcodes/inmemcache/cache"
)

func TestGet(t *testing.T) {
	cache := cache.NewCache[string, int]()
	cache.Set("a", 1)

	got, ok := cache.Get("a")
	if !ok {
		t.Fatalf("expected key %q to be found", "a")
	}
	if got != 1 {
		t.Errorf("got %d, want %d", got, 1)
	}

	_, ok = cache.Get("missing")
	if ok {
		t.Errorf("expected key %q to not be found", "missing")
	}
}

func TestSet(t *testing.T) {
	cache := cache.NewCache[string, int]()
	cache.Set("a", 1)
	cache.Set("a", 2)

	got, ok := cache.Get("a")
	if !ok {
		t.Fatalf("expected key %q to be found", "a")
	}
	if got != 2 {
		t.Errorf("got %d, want %d after overwrite", got, 2)
	}
}

func TestDelete(t *testing.T) {
	cache := cache.NewCache[string, int]()
	cache.Set("a", 1)
	cache.Delete("a")

	if _, ok := cache.Get("a"); ok {
		t.Errorf("expected key %q to be deleted", "a")
	}
}
