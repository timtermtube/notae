# notae is...
* JSON·map{} Based REST API·<del>Web Server</del> implementation in Golang, from Cheonan, South Korea.
* Simpler than others, but with both productivity and complexity for large REST API·Web Servers.
* Using Box-Component Structure
* Helping to code REST API simply
* <a href="https://pkg.go.dev/github.com/timtermtube/notae#section-documentation">Documentation is here.</a>

# Start in the Simplest way
```golang
// default address is: 0.0.0.0:8080
package main

import (
    "github.com/timtermtube/notae"
)

func main() {
    Comp := Notae.CreateComponent(Notae.ComponentOptions{
	Route:  "/",
	Title:  "A Component",
	Method: func(c *Component) {
            c.ModifyPlate("A", 1)
        }
    })
    Box := Notae.CreateBox("", "")
    Box.LinkComponent(Comp, 0)
    
    Box.LetGo()
}

```

