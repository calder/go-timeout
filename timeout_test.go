package timeout

import "fmt"
import "testing"
import "time"

func TestTimeout (T *testing.T) {
  start := time.Now()
  select {
    case <-Timeout(100 * time.Millisecond):
  }
  ms := time.Since(start).Nanoseconds() / 1000000
  if (ms < 100 || ms > 110) {
    fmt.Printf("Timeout took %dms.\n", ms)
    T.FailNow()
  }
}

func TestTwoTimeouts (T *testing.T) {
  select {
    case <-Timeout(100 * time.Millisecond):
    case <-Timeout(200 * time.Millisecond):
      fmt.Println("Longer timeout triggered first.")
      T.FailNow()
  }
}
