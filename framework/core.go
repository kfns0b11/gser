package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
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

// handler for method = GET
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error", err)
	}
}

// handler for method = POST
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("post router error", err)
	}
}

// handler for method = PUT
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("put router error", err)
	}
}

// handler for method = DELETE
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("delete router error", err)
	}
}

// http method wrap end
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// find handler for path
func (c *Core) FindRouteByRequest(req *http.Request) ControllerHandler {
	uri := req.URL.Path
	method := req.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)

	}
	return nil
}

func (c *Core) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext(rw, req)
	router := c.FindRouteByRequest(req)

	if router == nil {
		_ = ctx.Json(404, "not found")
		return
	}

	if err := router(ctx); err != nil {
		_ = ctx.Json(500, "inner error")
	}
}
