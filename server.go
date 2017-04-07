package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
