package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Selu", "TestSeluDefault", NewTestSeluDefault)
}

// NewTestSeluDefault version: 3.
func NewTestSeluDefault() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Selu",
		Title:  "TestSeluDefault",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x53, 0xa, 0xc, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x53, 0x65, 0x6c, 0x75, 0x12, 0x11, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6c, 0x75, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Selu",
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
				tensor.WithBacking([]float32{1.8534917, 0.4204456, 1.0283611, 2.3545086, 1.9622451, -1.0964665, 0.9982589, -0.2469415, -0.17241758, 0.43141627, 0.15134673, 1.5280067, 0.7996231, 0.12784407, 0.46636757, 0.35059196, 1.5698304, -0.32609507, 0.32894057, -1.009733, -1.6212339, 0.6867577, 0.908264, -0.9210997, 2.3848336, -1.3474979, 0.04807852, -0.300123, 1.6104927, 1.5438567, 0.16280343, 0.39733577, -1.0345254, -1.5155532, -0.5165983, 0.16427603, 1.2926676, 1.2633417, -0.56457984, -0.45866302, -1.1419833, -1.3331497, -1.4389311, 2.0496817, -0.70200115, -0.6236368, -1.2558014, 0.8169099, -1.4080441, -0.3369115, -1.0400617, 0.40651888, -0.7032181, -1.2182142, -0.048855502, 0.45004874, 0.06988971, 0.31780756, -0.8257883, -0.5348727}),
			),
		},
	}
}
