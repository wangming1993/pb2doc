package main

import (
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/pelletier/go-toml"
	"github.com/wangming1993/pb2doc/parser"
	"github.com/wangming1993/pb2doc/pb"
)

func main() {
	conf := "conf/local.toml"
	config, err := toml.LoadFile(conf)
	if err != nil {
		logrus.Fatalln(err)
	}

	package_prefix := config.GetDefault("package_prefix", "mairpc")
	parser.SetPrefix(package_prefix.(string))

	proto_base_path := config.GetDefault("proto_base_path", "./protos/proto/")
	parser.SetBasePath(proto_base_path.(string))

	proto_file := config.GetDefault("proto_file", "member/service.proto")
	protobuf := proto_file.(string)
	proto := &pb.Proto{}
	protos := proto.Initialize(filepath.Join(parser.GetBasePath(), protobuf))

	var messages []*pb.Message
	var services []*pb.Service
	for _, p := range protos {
		for _, m := range p.Messages {
			messages = append(messages, m.GetAll()...)
		}
		for _, s := range p.Services {
			services = append(services, s)
			s.WriteHtml()
		}
	}

	for _, m := range messages {
		m.WriteHtmlWithService(services)
	}
}
