package main

import (
	"os"
	"path/filepath"

	"github.com/docopt/docopt-go"
	"github.com/wangming1993/pb2doc/parser"
	"github.com/wangming1993/pb2doc/pb"
)

func main() {
	var (
		pkgPrefix        string
		protoPath        string
		distPath         string
		protoFilePattern string
	)
	usage := `
  Usage:
    pb2doc build <pkg> from <proto-dir> [--dist=<dir>]
    pb2doc serve <pkg> from <proto-dir>
    pb2doc -h | --help
    pb2doc --version

  Options:
    -h --help     Show this screen.
    --version     Show version.
    --dist=<dir>  HTML files containing folder [default: ./dist].`
	args, _ := docopt.Parse(usage, nil, true, "1.0", false)
	pkgPrefix = args["<pkg>"].(string)
	protoPath = args["<proto-dir>"].(string)
	distPath = args["--dist"].(string)
	protoFilePattern = ""

	parser.SetPrefix(pkgPrefix)
	parser.SetBasePath(protoPath)
	//TODO: use glob pattern
	fileName := filepath.Join(parser.GetBasePath(), protoFilePattern)

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

	// handle build logic
	if args["build"] == true {
		var messages []*pb.Message
		var services []*pb.Service
		var enums []*pb.Enum
		for _, p := range protos {
			for _, m := range p.Messages {
				messages = append(messages, m.GetAll()...)
				enums = append(enums, m.Enums...)
			}
			for _, s := range p.Services {
				services = append(services, s)
				s.WriteHtml(distPath)
			}
			enums = append(enums, p.Enums...)
		}

		for _, m := range messages {
			m.WriteHtmlWithService(distPath, services)
		}

		for _, e := range enums {
			e.WriteHtmlWithService(distPath, services)
		}
	}

	// handle serve logic
	if args["serve"] == true {
		// TODO: serve static html files to preview
	}
}
