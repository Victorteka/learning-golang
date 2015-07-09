Reference: [Request.FormFile](http://golang.org/pkg/net/http/#Request.FormFile)

```go
file, handler, err := r.FormFile("form_name_from_upload_tag")
```

handler.Filename now contains the original filename given by the browser.

You can then:

```go
data, err := ioutil.ReadAll(file)
```

This will load the file into "data". This is probably not ideal for big
uploads, but this is just an example, then you:

```go
err = ioutil.WriteFile(local_file_name, data, 0777)
```

to put it where you want.

Hope this helps.
