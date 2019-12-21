package main

import (
	"github.com/JettZheng/DesignPattern/composite"
)

func main()  {
	root := composite.NewComponent(composite.CompositeNode, "root")
	c1 := composite.NewComponent(composite.CompositeNode, "c1")
	c2 := composite.NewComponent(composite.CompositeNode, "c2")
	c3 := composite.NewComponent(composite.CompositeNode, "c3")

	l1 := composite.NewComponent(composite.LeafNode, "l1")
	l2 := composite.NewComponent(composite.LeafNode, "l2")
	l3 := composite.NewComponent(composite.LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}