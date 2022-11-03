package hago

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) insert(pattern string, parts []string, level int) {
	if len(parts) == level {
		n.pattern = pattern
		return
	}

	part := parts[level]
	var child *node

	for _, item := range n.children {
		if item.part == part || item.isWild {
			child = item
		}
	}

	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}

		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, level+1)
}

func (n *node) search(parts []string, level int) *node {
	if len(parts) == level || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[level]
	var children []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}

	for _, child := range children {
		result := child.search(parts, level+1)
		if result != nil {
			return result
		}
	}

	return nil
}
