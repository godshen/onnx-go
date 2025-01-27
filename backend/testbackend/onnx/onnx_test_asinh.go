package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Asinh", "TestAsinh", NewTestAsinh)
}

// NewTestAsinh version: 3.
func NewTestAsinh() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Asinh",
		Title:  "TestAsinh",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x4d, 0xa, 0xd, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x5, 0x41, 0x73, 0x69, 0x6e, 0x68, 0x12, 0xa, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x73, 0x69, 0x6e, 0x68, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Asinh",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.3328487, 0.3901813, 0.8662589, 1.546453, 1.3827868, -0.865215, 0.8456368, -0.15078518, -0.10303643, 0.39985773, 0.14355004, 1.1691281, 0.70195395, 0.12137677, 0.430447, 0.32777366, 1.1914744, -0.20374568, 0.3081669, -0.774355, -1.6667333, 0.6142558, 0.78219783, -0.6868674, 1.5581518, -1.1691804, 0.045742564, -0.18610767, 1.2128093, 1.1776454, 0.15433401, 0.36968425, -0.7997605, -1.4350144, -0.3412503, 0.1557189, 1.0352212, 1.0174958, -0.378243, -0.297878, -0.9152924, -1.1495624, -1.3039951, 1.4214023, -0.4898283, -0.42515022, -1.049338, 0.7149946, -1.2563262, -0.2111674, -0.8054935, 0.3778473, -0.4908553, -1.0035149, -0.028178498, 0.4162106, 0.06646827, 0.29803988, -0.59803224, -0.35522333}),
			),
		},
	}
}
