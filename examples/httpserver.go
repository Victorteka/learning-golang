package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
     `<DOCTYPE html>
      <html>
        <head>
            <title>Hello World</title>
        </head>
        <body>
            Hello World!
        </body>
      </html>`,
  )
}

func main() {
  http.HandleFunc("/hello", hello)

  // ---- static server -----
  // http.Handle(
  //   "/assets/",
  //   http.StripPrefix(
  //     "/assets/",
  //     http.FileServer(http.Dir("assets")),
  //   ),
  // )

  http.ListenAndServe(":6000", nil)
}
