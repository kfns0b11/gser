package framework

type IGroup interface {
	// implementing HttpMethod
	Get(string, ControllerHandler)
	Put(string, ControllerHandler)
	Post(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// implementing nest group
	Group(string) IGroup
}

type Group struct {
	core   *Core
	parent *Group
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// #region implementing IGroup interface

func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

func (g *Group) Group(subPrefix string) IGroup {
	childGroup := NewGroup(g.core, subPrefix)
	childGroup.parent = g
	return childGroup
}

// #endregion
