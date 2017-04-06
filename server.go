package sample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/{id:[0-9]+}", UserHandler).Methods("GET")
	r.Host("localhost")
	http.Handle("/", r)
}

type Reply struct {
	Message string `json:"message"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reply := Reply{
		Message: "id(" + id + ")を取得しました",
	}

	json, _ := json.Marshal(reply)

	fmt.Fprint(w, string(json))
}
