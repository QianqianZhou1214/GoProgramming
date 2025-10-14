package _select

import (
	"net/http"
)

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
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
