package cache_test

import (
	"testing"
	"time"

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

func TestGetWithTTL(t *testing.T) {
	c := cache.NewCache[string, string]()
	c.SetWithTTL("a", "hello", 50*time.Millisecond)

	got, ok := c.Get("a")
	if !ok {
		t.Fatalf("expected key %q to be found before ttl expires", "a")
	}
	if got != "hello" {
		t.Errorf("got %q, want %q", got, "hello")
	}

	time.Sleep(100 * time.Millisecond)

	if _, ok := c.Get("a"); ok {
		t.Errorf("expected key %q to be expired", "a")
	}
}

func TestSetNeverExpires(t *testing.T) {
	c := cache.NewCache[string, string]()
	c.Set("b", "world")

	time.Sleep(100 * time.Millisecond)

	got, ok := c.Get("b")
	if !ok {
		t.Fatalf("expected key %q without a ttl to never expire", "b")
	}
	if got != "world" {
		t.Errorf("got %q, want %q", got, "world")
	}
}

func TestSetWithTTLZeroIsImmediatelyExpired(t *testing.T) {
	c := cache.NewCache[string, string]()
	c.SetWithTTL("a", "hello", 0)

	if _, ok := c.Get("a"); ok {
		t.Errorf("expected key %q with zero ttl to be immediately expired", "a")
	}
}
