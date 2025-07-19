package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// The basic idea of this lesson is instead of making one function or unit responsible for all cancellations, we instead
// handle the cancellations inside the units themselves and propogate the errors using ctx.Err()
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, error := store.Fetch(r.Context())

		if error != nil {
			return
		}

		fmt.Fprint(w, res)
	}
}
