package main

type NodePath struct {
	n      *Element
	parent *NodePath
}

func GetFirstMoveInPath(end *NodePath) (int, int) {
	p := end
	for p.parent.parent != nil {
		p = p.parent
	}

	return p.n.x, p.n.y
}
