package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "welcome!")
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/todos", TodoIndex)
    router.GET("/todos/:todoId", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}
