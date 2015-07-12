package main

import (
    "net/http"
    "gopkg.in/unrolled/render.v1"
)

func main() {
    r := render.New(render.Options{})
    routes := http.NewServeMux()

    routes.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("Welcome, visit sub pages now."))
    })

    routes.HandleFunc("/data", func(res http.ResponseWriter, req *http.Request) {
        r.Data(res, http.StatusOK, []byte("Some binary data here."))
    })

    routes.HandleFunc("/json", func(res http.ResponseWriter, req *http.Request) {
        r.JSON(res, http.StatusOK, req)
    })

    routes.HandleFunc("/html", func(res http.ResponseWriter, req *http.Request) {
        // Assumes you have a template in ./templates called "example.tmpl"
        // $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
        r.HTML(res, http.StatusOK, "example", "Eugene")
    })

    http.ListenAndServe(":8080", routes)
}
