package main

import (
	"fmt"

	"github.com/brettmostert/fnple-go/ledger/http"
)

func main() {
	fmt.Println("Start the ledger service")
	http.Start()
}
