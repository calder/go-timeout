# Timeout
A teeny tiny timeout library for Go.

## Usage

```go
select {
case received := <- someChannel:
    // Do something
case <- timeout.Timeout(100 * time.Millisecond):
    // Abort
}
```
