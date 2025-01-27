package gorgonnx

import (
	"fmt"

	"github.com/godshen/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type reshape struct {
	toShape tensor.Shape
}

func init() {
	register("Reshape", newReshape)
}

func newReshape() operator {
	return &reshape{}
}

func (a *reshape) inferShape(requiredShape interface{}, targetShape tensor.Shape) error {
	var toShape tensor.Shape
	data := requiredShape
	if to, ok := data.(int64); ok {
		data = []int64{to}
	}
	if to, ok := data.([]int64); ok {
		childShape := make([]int, len(to))
		copy(childShape, targetShape)
		toShape = make([]int, len(to))
		dimSize := 1
		for i := 0; i < len(childShape); i++ {
			dimSize *= childShape[i]
		}
		for i := 0; i < len(to); i++ {
			toShape[i] = int(to[i])
		}
		for i := 0; i < len(toShape); i++ {
			if toShape[i] == 0 {
				toShape[i] = childShape[i]
			}
		}
		actualSize := 1
		for i := 0; i < len(toShape); i++ {
			actualSize *= toShape[i]
		}
		for i := 0; i < len(toShape); i++ {
			if toShape[i] == -1 {
				toShape[i] = dimSize / actualSize
				if toShape[i] == 0 {
					toShape = append(toShape[:i], toShape[i+1:]...)
				}
			}
		}
	} else {
		return fmt.Errorf("Cannot reshape, bad output shape %#v", requiredShape)
	}
	a.toShape = toShape
	return nil
}

func (a *reshape) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}

	err = a.inferShape(children[1].gorgoniaNode.Value().Data(), children[0].gorgoniaNode.Shape())
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Reshape(children[0].gorgoniaNode, a.toShape)

	return err
}

func (a *reshape) init(o onnx.Operation) error {
	return nil
}
