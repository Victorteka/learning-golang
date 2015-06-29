# Go: The Good Parts

Go is an opinionated programming language initially developed by Google in 2007. It is a statically typed language with a simple syntax, resembling C or JavaScript. It features garbage collection, type safety and large standard library.
---

Go can be statically compiled into a single executable binary, which could target a large number of operating systems (from Linux and Windows to Plan 9) and processors (i386, amd64 and ARM).

In HappyPancake project we found that Golang was a good fit for developing event-driven backend services. Other contenders included Scala, C#, Haskell and Erlang.
Here is how code can look like in Golang:

```go
func (m *ContactsModule) handleList(r *ApiRequest) ApiResponse {
    return NewObjectResponse(&ContactListModel{})
}

type ContactsModule struct{}

func (m *ContactsModule) Register(r Runtime) {
    r.HandleApi("GET", "/contacts", m.handleList)
}
```
