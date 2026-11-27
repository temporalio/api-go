package recursive

type Payload struct{}

type Node struct {
	Child   *Node
	Payload *Payload
}

func (n *Node) GetChild() *Node {
	return n.Child
}

func (n *Node) GetPayload() *Payload {
	return n.Payload
}
