# ghoST
ghoST is..
* JSON·map{} Based REST API·<del>Web Server</del> implementation in Golang
* Simpler than others, but with both productivity and complexity for large REST API·Web Servers.
* Using Box-Component Structure
* Helping to code REST API simply
* <a href="https://github.com/timtermtube/goST/wiki">also Documentations!</a>
* <a href="https://pkg.go.dev/github.com/timtermtube/goST@main#section-documentation">API References</a>

# Start in the Simplest way
```golang
// default address is: 0.0.0.0:8080
package main

import (
    "github.com/timtermtube/ghoST"
)

func main() {
    Comp := ghoST.CreateComponent(ghoST.ComponentOptions{
	Route:  "/",
	Title:  "A Component",
	Method: func(c *Component) {
            c.ModifyPlate("A", 1)
        }
    })
    Box := ghoST.CreateBox("", "")
    Box.LinkComponent(Comp, 0)
    
    Box.LetGo()
}

```

