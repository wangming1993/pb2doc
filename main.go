package main

import (
	"path/filepath"

	"os"

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

	fileName := filepath.Join(parser.GetBasePath(), protobuf)
	var protos []*pb.Proto

	if parser.IsDir(fileName) {
		filepath.Walk(fileName, func(path string, info os.FileInfo, err error) error {
			if parser.IsDir(path) {
				return err
			}
			if !parser.IsProtoFile(path) {
				return err
			}
			proto := &pb.Proto{}
			ps := proto.Initialize(path)
			protos = append(protos, ps...)
			return err
		})
	} else {
		proto := &pb.Proto{}
		protos = proto.Initialize(fileName)
	}

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
