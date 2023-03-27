package core

// 树节点
type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node // 子节点
	isWild   bool    //是否精确匹配:part含有 : || true 时为空
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
// func (n *node) insert(pattern string, parts []string, height int) {
// 	if len(parts) == height {
// 		n.pattern = pattern
// 		return
// 	}

// 	part := parts[height]
// 	child := n.matchChild(part)
// 	if child == nil {
// 		child := &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
// 		n.children = append(n.children, child)
// 	}
// }

// 搜索节点
// func (n *node)
