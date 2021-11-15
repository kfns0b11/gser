package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = newTree()
	router["POST"] = newTree()
	router["PUT"] = newTree()
	router["DELETE"] = newTree()

	return &Core{
		router: router,
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// handler for method = GET
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error", err)
	}
}

// handler for method = POST
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("post router error", err)
	}
}

// handler for method = PUT
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("put router error", err)
	}
}

// handler for method = DELETE
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("delete router error", err)
	}
}

// http method wrap end
func (c *Core) Group(prefix string) IGroup {
	return newGroup(c, prefix)
}

func (c *Core) FindRouteNodeByRequest(req *http.Request) *node {
	uri := req.URL.Path
	method := req.Method
	methodUpper := strings.ToUpper(method)

	if methodHandlers, ok := c.router[methodUpper]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext(rw, req)
	node := c.FindRouteNodeByRequest(req)
	if node == nil {
		ctx.SetStatus(404).Json("not found")
		return
	}

	ctx.SetHandlers(node.handlers)

	params := node.parseParamsFromEndNode(req.URL.Path)
	ctx.SetParams(params)

	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}
}
