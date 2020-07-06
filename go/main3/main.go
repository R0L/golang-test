package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	// ...
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

func main() {

	fmt.Println()

}
