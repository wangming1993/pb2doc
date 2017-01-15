package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/wangming1993/pb2doc/parser"
	"github.com/wangming1993/pb2doc/pb"
)

func main() {
	logrus.Println("init project...")
	var protobuf = "./protos/test.proto"

	proto := &pb.Proto{}
	proto.Initialize(protobuf)

	cm := "  // The name of person"
	fmt.Println(parser.IsSingleComment(cm))
}
