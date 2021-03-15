```export CGO_CFLAGS="-g -O2 -Wno-return-local-addr"```
```go build --tags "sqlite_fts5" && ./backend```