package http

import (
	"fmt"
	"net/http"

	"github.com/brettmostert/fnple-go/internal/ctx"
)

func handleAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// acc := accounts.NewAccount{
		// 	CoalationId: uuid.New().String(),
		// 	Description: "123",
		// }
		// newAcc := accounts.CreateAccount(s.ctx, acc)
		// b, _ := json.MarshalIndent(newAcc, "", "  ")
		// print(string(b))
		fmt.Fprintf(w, "Hello, %q %v", ctx.DB(r.Context()), ctx.RequestID(r.Context()))
	}
}
