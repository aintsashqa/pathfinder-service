package types

type Node struct {
	Position Point
	Weight   float64
	Distance float64
	Previous *Node
}

func (n *Node) Cost() float64 {
	return n.Weight + n.Distance
}
