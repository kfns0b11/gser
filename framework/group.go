package framework

type IGroup interface {
	// implementing HttpMethod
	Get(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	// implementing nest group
	Group(string) IGroup

	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core   *Core
	parent *Group
	prefix string

	middlewares []ControllerHandler
}

func newGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,

		middlewares: []ControllerHandler{},
	}
}

// #region implementing IGroup interface

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.middlewares, handlers...)
	g.core.Get(uri, allHandlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.middlewares, handlers...)
	g.core.Post(uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.middlewares, handlers...)
	g.core.Put(uri, allHandlers...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.middlewares, handlers...)
	g.core.Delete(uri, allHandlers...)
}

func (g *Group) Group(subPrefix string) IGroup {
	childGroup := newGroup(g.core, subPrefix)
	childGroup.parent = g
	return childGroup
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// #endregion

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}
