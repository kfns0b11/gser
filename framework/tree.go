package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast   bool
	segment  string
	handlers []ControllerHandler

	childs []*node
	parent *node
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
		parent:  nil,
	}
}

func newTree() *Tree {
	root := newNode()
	return &Tree{root}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// Filtering child nodes of n which match the segment.
// If the segment is wild, all child nodes are matched.
// If the child node segment is wild, matched.
// If the child node segment equals the segment, matched.
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	nodes := n.filterChildNodes(segment)

	if len(nodes) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, v := range nodes {
			if v.isLast {
				return v
			}
		}

		return nil
	}

	for _, v := range nodes {
		vmatch := v.matchNode(segments[1])
		if vmatch != nil {
			return vmatch
		}
	}

	return nil
}

func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")

	cnt := len(segments)
	cur := n

	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}

		if isWildSegment(cur.segment) {
			ret[cur.segment[1:]] = segments[i]
		}

		cur = cur.parent
	}
	return ret
}

func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	lenth := len(segments)

	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		isLast := index == lenth-1
		var objNode *node

		childNodes := n.filterChildNodes(segment)

		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}
