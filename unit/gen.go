package unit

//go:generate protoc -I../../../../ -I ../../../../github.com/gogo/protobuf/protobuf -I. --gogo_out=. unit.proto
//go:generate ffjson source_unit.go
