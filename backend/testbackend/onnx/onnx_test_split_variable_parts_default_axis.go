package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Split", "TestSplitVariablePartsDefaultAxis", NewTestSplitVariablePartsDefaultAxis)
}

// NewTestSplitVariablePartsDefaultAxis version: 3.
func NewTestSplitVariablePartsDefaultAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Split",
		Title:  "TestSplitVariablePartsDefaultAxis",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa1, 0x1, 0xa, 0x32, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x22, 0x5, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x2a, 0xe, 0xa, 0x5, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x40, 0x2, 0x40, 0x4, 0xa0, 0x1, 0x7, 0x12, 0x26, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0x13, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"input"},
		     Output:    []string{"output_1", "output_2"},
		     Name:      "",
		     OpType:    "Split",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000119500)(name:"split" type:INTS ints:2 ints:4 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 2}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]float32{3, 4, 5, 6}),
			),
		},
	}
}
