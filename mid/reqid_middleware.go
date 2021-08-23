package mid

import (
	"net/http"

	"gitlab.com/gocor/corctx"
)

// RequestIDHandler is a middleware that injects a request ID into the context of each
// request.
func RequestIDHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := corctx.WithHTTPRequest(r)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
