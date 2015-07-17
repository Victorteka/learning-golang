package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    // "strings"
    "crypto/md5"
    // "encoding/json"
    "gopkg.in/unrolled/render.v1"
    "io"
    "os"
    "strconv"
    "time"
)

// upload logic
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.html")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)

        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        // fmt.Fprintf(w, "%v", handler.Header)
        Reponse := render.New(render.Options{})
        Reponse.JSON(w, http.StatusOK, r)

        f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()

        io.Copy(f, file)
    }
}

func main() {
    http.HandleFunc("/", uploadHandler) // setting router rule

    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
