package gorgonnx

import (
	"github.com/godshen/onnx-go"
)

// dropout is a void function that does nothing because we are supposed to be in test mode

type dropout struct{}

func init() {
	register("Dropout", newDropout)
}

func newDropout() operator {
	return &dropout{}
}

func (a *dropout) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}
	n.gorgoniaNode = children[0].gorgoniaNode
	return err
}

func (a *dropout) init(o onnx.Operation) error {
	return nil
}
