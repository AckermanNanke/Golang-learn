package core

// 树节点
type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node // 子节点
	isWild   bool    //是否精确匹配:part含有 : || true 时为空
}

// 匹配节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}
