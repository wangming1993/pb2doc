package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/wangming1993/pb2doc/parser"
	"github.com/wangming1993/pb2doc/pb"
)

func main() {
	logrus.Println("init project...")
	var protobuf = "test.proto"

	parser.SetBasePath("./protos/")
	proto := &pb.Proto{}
	proto.Initialize(parser.ProtoPath + protobuf)
}
