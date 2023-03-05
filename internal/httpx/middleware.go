package httpx

import (
	"log"
	"net/http"

	"github.com/brettmostert/fnple-go/internal/auth"
	"github.com/brettmostert/fnple-go/internal/ctx"
	"github.com/google/uuid"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func CompileMiddleware(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func WithRequestLogging(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s: %s", r.Method, r.RequestURI, r.Body)
		h.ServeHTTP(w, r)
	})
}

func WithDB(db string) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ctx.WithDB(r.Context(), db)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth.Auth()
		h.ServeHTTP(w, r)
	})
}

func WithRequestID(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := ctx.WithRequestID(r.Context(), uuid.NewString())
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
