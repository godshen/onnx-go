package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Transpose", "TestTransposeAllPermutations2", NewTestTransposeAllPermutations2)
}

// NewTestTransposeAllPermutations2 version: 3.
func NewTestTransposeAllPermutations2() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Transpose",
		Title:  "TestTransposeAllPermutations2",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x91, 0x1, 0xa, 0x2e, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0xa, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x22, 0x9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x2a, 0xf, 0xa, 0x4, 0x70, 0x65, 0x72, 0x6d, 0x40, 0x1, 0x40, 0x0, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x21, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x32, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x20, 0xa, 0xa, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"transposed"},
		     Name:      "",
		     OpType:    "Transpose",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000126100)(name:"perm" type:INTS ints:1 ints:0 ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 2, 4),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.9786183, 0.7991586, 0.46147937, 0.7805292}),
			),
		},
	}
}
