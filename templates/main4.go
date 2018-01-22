package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "encoding/json"
    "time"

    "github.com/julienschmidt/httprouter"
)

type Todo struct {
    Name            string      `json:"name"`
    Completed       bool        `json:"completed"`
    Due             time.Time   `json:"due"`
}

type Todos []Todo

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    todos := Todos {
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    json.NewEncoder(w).Encode(todos)
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
