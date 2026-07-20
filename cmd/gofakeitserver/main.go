package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

var port string
var faker *gofakeit.Faker

func init() {
	flag.StringVar(&port, "port", "8080", "server port")
	faker = gofakeit.New(0)
}

func main() {
	flag.Parse()

	gofakeit.Seed(0)

	mux := http.NewServeMux()
	routes(mux)

	fmt.Println("Running on port " + port)

	http.ListenAndServe(":"+port, mux)
}

func routes(mux *http.ServeMux) { _ = "STUB: not implemented"; return }

func favicon(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func list(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func lookup(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func lookupGet(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func lookupPost(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func getInfoFromPath(r *http.Request) (*gofakeit.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeResponse(v any) []byte { _ = "STUB: not implemented"; return nil }

func ok(w http.ResponseWriter, data any) { _ = "STUB: not implemented"; return }

func badrequest(w http.ResponseWriter, msg string) { _ = "STUB: not implemented"; return }

func notfound(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
