package backend

import "github.com/godshen/onnx-go"

// ComputationBackend is a backend that can run the graph
type ComputationBackend interface {
	onnx.Backend
	Run() error
}
