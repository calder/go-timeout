package timeout

import "time"

func Timeout (timeout time.Duration) chan bool {
    channel := make(chan bool)
    go func () {
        time.Sleep(timeout)
        channel <- true
    }()
    return channel
}
