package core

import "strings"

// 树节点
type node struct {
	pattern  string  //待匹配路由的全地址 例：/a/:b/c
	part     string  //路由中的一部分 例：:b
	children []*node // 子节点
	isWild   bool    //是否精确匹配	part含有 : 或 * 时为 true
}

// 第一个匹配节点	插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点	查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入节点
func (n *node) insert(pattern string, parts []string, index int) {
	if len(parts) == index {
		n.pattern = pattern
		return
	}

	part := parts[index]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, index+1)
}

// 搜索节点
func (n *node) search(parts []string, index int) *node {
	if len(parts) == index || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[index]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, index+1)
		if result != nil {
			return result
		}
	}
	return nil
}
