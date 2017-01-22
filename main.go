package main

import (
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/wangming1993/pb2doc/parser"
	"github.com/wangming1993/pb2doc/pb"
)

func main() {
	logrus.Println("init project...")
	protobuf := "member/service.proto"

	parser.SetBasePath("./protos/proto/")
	parser.SetPrefix("mairpc")

	proto := &pb.Proto{}
	proto.Initialize(filepath.Join(parser.GetBasePath(), protobuf))
}
