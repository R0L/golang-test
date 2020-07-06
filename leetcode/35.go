package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	var (
		rNode = new(Node)
		nMap  = make(map[*Node]*Node)
	)
	for node := head; node != nil; node = node.Next {
		n := copyNode(nMap, node)

	}
	return rNode
}

func copyNode(nMap map[*Node]*Node, head *Node) *Node {
	if nil == head {
		return nil
	}
	dNode, ok := nMap[head]
	if !ok {
		dNode = new(Node)
		nMap[head] = dNode
	}
	dNode.Val = head.Val
	if nil != head.Random {
		rNode := new(Node)
		nMap[head.Random] = rNode
		dNode.Random = rNode
	}
	return dNode
}
