package main

import (
	"fmt"
	"strings"
)

type handler struct {
	name, pattern  string
	handler        interface{}
	method, filter string
	isRegex, depth int
	meta           map[string]interface{}
}

type path struct {
	handler
	node map[rune]map[string]*path
}

func getCharCodeAt(s string, n int) rune {
	for i, r := range s {
		if i == n {
			return r
		}
	}
	return 0
}

func transverseNode(branch *path, part string) (*path, bool) {

	// get the charcode of the 1st character of the route part
	charcode := getCharCodeAt(part, 0)

	// now check if it's node exists
	if pathmap, ok := branch.node[charcode]; ok {

		// if the node exists, then does it's path details exist
		if path, ok := pathmap[part]; ok {
			// if yes return it
			return path, true
		}

		// it's path details do not exist
		return branch, false
	}

	// If not return false, with the branch itself
	// false bool will aid in breaking the loop
	return branch, false
}

func (p *path) getNode(pattern string) *path {

	routeParts := strings.Split(pattern[1:], "/")
	fmt.Println(routeParts)

	var branchpath *path

	var pathFound bool
	for _, part := range routeParts {

		// If a path is not assigned to branch path, give it the default one
		if branchpath == nil {
			// root path/handler
			branchpath = p
		}

		// If a path is returned then it is the nested one
		// and fed back to get the next nested branch
		// with the new route part
		branchpath, pathFound = transverseNode(branchpath, part)

		// If no nested path is found, break it there is no match
		if !pathFound {
			// nothing found
			break
		}
	}

	if branchpath == nil {
		return nil
	}
	return branchpath
}

func main() {
	r := new(path)
	h := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/user handler here",
	}

	h2 := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/{name} handler for /users here",
	}

	h3 := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/users handler here",
	}

	h4 := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/images handler for the route /user/{name}",
	}

	h5 := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/{id} handler for the route /user/{name}/images",
	}

	h6 := handler{
		name:    "{name}",
		pattern: "/user/{name}",
		handler: "/{id} handler for the route /user/{name}/images",
	}

	r.node = map[rune]map[string]*path{
		117: {
			"user": &path{
				handler: h,
				node: map[rune]map[string]*path{
					123: {
						"{name}": &path{
							handler: h2,
							node: map[rune]map[string]*path{
								105: {
									"images": &path{
										handler: h4,
										node: map[rune]map[string]*path{
											123: {
												"{id}": &path{
													handler: h5,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"users": &path{
				handler: h3,
				node: map[rune]map[string]*path{
					99: {
						"{categories}": &path{
							handler: h6,
						},
					},
				},
			},
		},
	}

	b := r.getNode("/users/{categories}")
	c := r.getNode("/user/{name}/images/{id}")

	fmt.Printf("%v\n", b.handler.handler)
	fmt.Printf("%v\n", c.handler.handler)
}
