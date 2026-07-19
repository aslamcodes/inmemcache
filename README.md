# inmemcache

An in-memory cache built from scratch in Go, purely as a vehicle to learn Go. The learning is the point, not the cache.

## End product

Three separate deliverables, layered on top of each other:

1. **`cache` library** — a standalone, importable, generic `Cache[K comparable, V any]` engine (TTL, LRU eviction, concurrency-safe, sharded). Knows nothing about networking or servers; usable in-process by any Go program, including this project's own server.
2. **`cacheserver`** — a binary that wraps the library (`cache.New[string, []byte]()`), adds a wire protocol (TCP or HTTP), persistence (snapshot to disk / reload on startup), and exposes stats (hits/misses/evictions).
3. **`cachecli`** — a standalone CLI binary that speaks the server's protocol as a client, for interactive use (`set`, `get`, `del`, etc.).

Each is a distinct piece of learning: generics + concurrency (library), networking + persistence (server), protocol/client design + CLI ergonomics (cli).

## Feature checklist (no schedule — build in whatever order makes sense, driven by what's next logically)

**1. `cache` library**
- [x] Generics — `Cache[K comparable, V any]`
- [x] `Cache` struct, `Get`/`Set`/`Delete` on a plain map, table-driven tests
- [x] Concurrency-safe (`sync.RWMutex`) — reproduce and understand the unsafe-map crash first
- [ ] TTL — `SetWithTTL`, lazy expiry on read, background sweeper goroutine with graceful shutdown (`context` + `sync.WaitGroup`)
- [ ] Eviction policy — capacity limit + LRU (doubly linked list + map) behind an interface so policies are pluggable
- [ ] Sharding to reduce lock contention; benchmark before/after with `go test -bench` / `pprof`

**2. `cacheserver`**
- [ ] Network layer — expose the library over TCP (custom line protocol) or HTTP; instantiates `cache.New[string, []byte]()`
- [ ] Persistence — periodic snapshot to disk, reload on startup, save-on-exit via `os.Signal`
- [ ] Observability — hit/miss/eviction stats, structured logging

**3. `cachecli`**
- [ ] Protocol client (dial the server, send/parse commands) — likely its own small package, reusable by other Go clients too
- [ ] Interactive REPL (`set`/`get`/`del`/`stats`) and/or one-shot command mode

Some ordering is load-bearing (concurrency-safety before TTL's background goroutine; library before server wraps it; server before CLI has anything to talk to), but there's no calendar attached — work feature to feature, not day to day.

## Project layout (once past the single-file stage)

```
cache/                     // 1. the library — public, importable, no networking
cmd/cacheserver/main.go    // 2. the server
cmd/cachecli/main.go       // 3. the cli
internal/proto/            // wire protocol shared by server + cli client
```
