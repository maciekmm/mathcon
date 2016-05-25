package mathcon

import (
	"net/http"

	"golang.org/x/net/context"
)

func contestHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	// if r.Method == http.MethodPost {
	// 	submitQuestion(ctx, w, r)
	// } else if r.Method == http.MethodGet {

	// }
	// fmt.Fprint(w, "Hello, world!")
}

func submitQuestion(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}
