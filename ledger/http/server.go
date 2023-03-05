package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brettmostert/fnple-go/internal/ctx"
	"github.com/brettmostert/fnple-go/internal/httpx"
)

func Start() {
	fmt.Println("Starting http server")

	// TODO: Get info from configuration
	ctx := &ctx.AppContext{
		DB: "conn",
	}

	fmt.Println(ctx.DB)

	mw := []httpx.Middleware{
		httpx.WithRequestID,
		httpx.WithRequestLogging,
		httpx.WithDB("moodb2"),
		httpx.WithAuth,
	}

	router := http.NewServeMux()
	router.HandleFunc("/accounts", httpx.CompileMiddleware(handleAccount(), mw))

	server := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	_ = server.ListenAndServe()
}
