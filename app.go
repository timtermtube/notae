package Notae

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// For Test

type CompMethod func(c *Component)

type Request struct {
	Type   string
	Query  url.Values
	Header map[string]interface{}
	Data   string
}

type Component struct {
	Id             int
	Title          string
	Route          string
	HttpCode       int
	Method         CompMethod
	Request        Request
	ResponseHeader map[string]string
	ResponsePlate  map[string]interface{}
}

type CompOptions struct {
	Title  string
	Route  string
	Method CompMethod
}

type Box struct {
	CallSign         string
	AppTitle         string
	Address          string
	Components       []Component
	FunctionalRouter http.ServeMux
}

type BoxHttpsOptions struct {
}

func CreateBox(AppTitle string, Address string) Box {
	if Address == "" {
		Address = ":8080"
	}
	if AppTitle == "" {
		AppTitle = "Default Notae Application"
	}

	CreatedBox := Box{
		AppTitle:   AppTitle,
		Address:    Address,
		Components: []Component{},
		CallSign:   fmt.Sprintf("<Box:%s[%v]>", AppTitle, Address),
	}
	return CreatedBox
}

func (Box *Box) Run() {
	fmt.Printf("Your Notae Application: <%s> is started to working on HTTP, <%s>\n", Box.AppTitle, Box.Address)
	err := http.ListenAndServe(Box.Address, &Box.FunctionalRouter)
	if err != nil {
		log.Fatal(fmt.Sprintf("Server Runtime Internal Error: <%s>\n", err))
	}
}

func (Box *Box) LinkComponent(Component Component, id int) {
	Component.Id = id
	for _, Cpnt := range Box.Components {
		if Cpnt.Id == Component.Id {
			log.Fatal(fmt.Sprintf("Component ID: <%v> exists at Component: <%s>\n", Cpnt.Id, Cpnt.Title))
		} else if Cpnt.Route == Component.Route {
			log.Fatal(fmt.Sprintf("The requested Route: <%s> is already occupied by Component ID: <%v>\n", Component.Route, Cpnt.Id))
		}
	}
	Processor := func(res http.ResponseWriter, req *http.Request) {
		dialogue := fmt.Sprintf("Method: <%s> to Component of Box: %s: <%v:%v> from <%v> ",
			req.Method, Box.CallSign, Component.Title, Component.Id, strings.Split(req.RemoteAddr, ":")[0])
		fmt.Println(dialogue)
		TempChan := make(chan string)
		go func(channel chan string) {
			TempByte := new(bytes.Buffer)
			TempByte.ReadFrom(req.Body)
			Component.HttpCode = 200
			Component.ResponseHeader = map[string]string{"Content-Type": "application/json; charset=utf-8"}
			Component.ResponsePlate = make(map[string]interface{})
			Component.Request = Request{
				Type:   req.Method,
				Query:  req.URL.Query(),
				Header: map[string]interface{}{},
				Data:   TempByte.String(),
			}
			for k, v := range req.Header {
				Component.Request.Header[k] = v
			}
			Component.Method(&Component)
			Component.ResponsePlate["code"] = Component.HttpCode
			jsonified, _ := json.Marshal(Component.ResponsePlate)
			for k, v := range Component.ResponseHeader {
				res.Header().Add(k, v)
			}
			res.WriteHeader(Component.HttpCode)
			channel <- string(jsonified)
		}(TempChan)

		Value := <-TempChan
		res.Write([]byte(Value))
	}

	Box.Components = append(Box.Components, Component)
	Box.FunctionalRouter.HandleFunc(Component.Route, Processor)
}

func (Box *Box) DetachComponent(ComponentID int) {
	RefreshedComponent := []Component{}

	for _, Cpnt := range Box.Components {
		if Cpnt.Id != ComponentID {
			RefreshedComponent = append(RefreshedComponent, Cpnt)
		}
	}

	Box.Components = RefreshedComponent
}

func (Box *Box) LetGo() {
	fmt.Printf("Box Title <%s>: Started working on the Address: <%v>. The callsign is: %s.\n", Box.AppTitle, Box.Address, Box.CallSign)
	e := http.ListenAndServe(Box.Address, &Box.FunctionalRouter)

	if e != nil {
		ErrorDialogue := fmt.Sprintf("Box Callsign %s: Caught Start Error: <%v>\n", Box.CallSign, e)
		log.Fatal(ErrorDialogue)
	}
}

func CreateComponent(Options CompOptions) Component {
	if Options.Route == "" {
		log.Fatal(fmt.Sprintf("You must link the route to Component.\n"))
	} else if string(Options.Route[0]) != "/" {
		Options.Route = fmt.Sprintf("/%s", Options.Route)
	}

	ReturnComponent := Component{
		Title:          Options.Title,
		Route:          Options.Route,
		HttpCode:       200,
		Method:         Options.Method,
		ResponseHeader: map[string]string{"Content-Type": "application/json; charset=utf-8"},
		ResponsePlate:  make(map[string]interface{}),
	}

	return ReturnComponent
}

func (c *Component) ModifyPlate(key string, data interface{}) {
	c.ResponsePlate[key] = data
}

func (c *Component) ModifyHttpHeader(name string, value string) {
	c.ResponseHeader[name] = value
}
