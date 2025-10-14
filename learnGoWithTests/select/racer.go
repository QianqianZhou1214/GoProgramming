package _select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeOut time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeOut):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{}) // always make channels, do not do var ch because it's nil and cannot assign value by <-.
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
