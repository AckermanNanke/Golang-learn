package core

type nodeType uint8

type node struct {
	path      string
	indices   string
	wildChild bool
	nType     nodeType
	priority  uint32
	children  []*node // 子节点：数组末尾最多有1个param样式的节点
	handlers  HandlersChain
	fullPath  string
}
type methodTree struct {
	method string
	root   *node
}
type methodTrees []methodTree
