package mathcon

import (
	"fmt"
	"net/http"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
