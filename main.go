package main

import (
	"fmt"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
	l5g "github.com/neocortical/log5go"
)

var mc *memcache.Client
var log = l5g.Logger(l5g.LogAll)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/test", TestHandler)
	http.Handle("/", r)

	mc = memcache.New("172.17.0.6:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	log.Fatal("http error: %v", http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Client: %+v\n", mc)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Test val: %s\n", string(it.Value))
}
