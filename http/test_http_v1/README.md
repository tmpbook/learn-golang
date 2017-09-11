## Fix the race

If your code has data races, all bets are off and you're just waiting for a crash. The runtime promises nothing if you have a data race.

Multiple options:

- use channels
- use a Mutex
- use atomic

### Mutex

```go
var visitors struct {
    sync.Mutex
    n int
}
...
func foo() {
    ...
    visitors.Lock()
    visitors.n++
    yourVisitorNumber := visitors.n
    visitors.Unlock()
```
[used here](./demo_test.go#L11-L14)

### Atomic

```go
var visitors int64 // must be accessed atomically
...
func foo() {
    ...
    visitNum := atomic.AddInt64(&visitors, 1)
```