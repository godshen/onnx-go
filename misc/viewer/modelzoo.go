package main

import (
	"fmt"
	"log"
	"os"

	"github.com/godshen/onnx-go/internal/onnx/ir"
	"github.com/gogo/protobuf/proto"
	"github.com/kr/pretty"
)

func main() {
	b, err := os.ReadFile("/Users/olivier.wulveryck/Documents/squeezenet/model.onnx")
	if err != nil {
		log.Fatal(err)
	}
	model := &ir.ModelProto{}
	err = proto.Unmarshal(b, model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%# v", pretty.Formatter(model))

}
