package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Softmax", "TestSoftmaxLargeNumber", NewTestSoftmaxLargeNumber)
}

// NewTestSoftmaxLargeNumber version: 3.
func NewTestSoftmaxLargeNumber() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Softmax",
		Title:  "TestSoftmaxLargeNumber",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x56, 0xa, 0xf, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Softmax",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 4),
				tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 4),
				tensor.WithBacking([]float32{0.032058604, 0.08714432, 0.23688284, 0.6439143, 0.032058604, 0.08714432, 0.23688284, 0.6439143}),
			),
		},
	}
}
