package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
		/*
			ctx := r.Context()

			data := make(chan string, 1)

			go func() {
				data <- store.Fetch()
			}()
			select {
			case d := <-data:
				fmt.Fprint(w, d)
			case <-ctx.Done():
				store.Cancel()
			}

		*/

	}
}

/*
client  ----->  Server(store)
                  |
                  | starts goroutine --> store.Fetch()
                  |
                  | listen for two things:
                  |    1. data ready  --> send response
                  |    2. ctx.Done()  --> cancel store

*/
