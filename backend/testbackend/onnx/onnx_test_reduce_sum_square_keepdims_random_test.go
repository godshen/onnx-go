package onnxtest

import (
	"testing"

	"github.com/godshen/onnx-go/internal/onnx/ir"
	"github.com/gogo/protobuf/proto"
)

func TestNewTestReduceSumSquareKeepdimsRandom(t *testing.T) {
	mytest := NewTestReduceSumSquareKeepdimsRandom()
	var model ir.ModelProto
	err := proto.Unmarshal(mytest.ModelB, &model)
	if err != nil {
		t.Fatal(err)
	}
	if model.Graph == nil {
		t.Fatal("graph is nil")
	}
	if len(model.Graph.Input) != len(mytest.Input) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Input), len(mytest.Input))
	}
	if len(model.Graph.Output) != len(mytest.ExpectedOutput) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Output), len(mytest.ExpectedOutput))
	}
}
