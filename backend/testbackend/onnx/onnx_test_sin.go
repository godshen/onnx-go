package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/godshen/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Sin", "TestSin", NewTestSin)
}

// NewTestSin version: 3.
func NewTestSin() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Sin",
		Title:  "TestSin",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x49, 0xa, 0xb, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x3, 0x53, 0x69, 0x6e, 0x12, 0x8, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Sin",
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
				tensor.WithBacking([]float32{0.9813841, 0.38956314, 0.82979375, 0.78376156, 0.95628846, -0.828978, 0.81346697, -0.15077996, -0.103035666, 0.39915818, 0.14354597, 0.9932189, 0.68967324, 0.12137501, 0.42943156, 0.3275169, 0.9970587, -0.2037221, 0.30797866, -0.75397724, -0.5551996, 0.6080631, 0.76072943, -0.67588514, 0.7655128, -0.9932296, 0.04574255, -0.18609269, 0.9992774, 0.99485964, 0.15432815, 0.36921346, -0.77567613, -0.91712075, -0.3409358, 0.15571275, 0.9425859, 0.9328988, -0.3777146, -0.29771933, -0.8667023, -0.98865443, -0.99083745, 0.92867243, -0.48787367, -0.4241964, -0.94986236, 0.7014931, -0.9990713, -0.2111392, -0.7805008, 0.3773217, -0.48887977, -0.92484665, -0.028178498, 0.41535395, 0.06646818, 0.2978808, -0.59263164, -0.35483837}),
			),
		},
	}
}
