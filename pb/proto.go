package pb

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/wangming1993/pb2doc/parser"
)

type Proto struct {
	Messages []*Message
	Services []*Service
	Comment  string
	Package  string
	content  []string
	Start    int // the message start int, excludes syntax,package and imports line
	Syntax   string
	Imports  []string
	Path     string // base folder of proto file
}

func (p *Proto) Initialize(file string) []*Proto {
	lines := parser.ReadFile(file)
	p.content = lines
	supported := p.IsSupported()
	if !supported {
		panic(errors.New("Unsupported proto syntax... file:" + file))
	}
	p.Syntax = p.syntax()
	p.Start += 1 //excludes syntax line

	p.initPackage()
	p.initImports()

	RegisterProto(p.Package, parser.FileName(file))

	var protos []*Proto
	ps := p.ResolveImports()
	if ps != nil {
		protos = append(protos, ps...)
	}

	p.Parse()
	protos = append(protos, p)
	//fmt.Println(p.JSON())
	return protos
}

func (p *Proto) syntax() string {
	syntax_pattern := "^syntax\\s?=\\s?\"(.+)\"\\s?;"
	cp := regexp.MustCompile(syntax_pattern)
	matches := cp.FindStringSubmatch(p.content[0])
	if len(matches) > 1 {
		return matches[1]
	}
	return PROTO_SYNTAX_UNKNOWN
}

func (p *Proto) initPackage() {
	package_pattern := "^package\\s+(.+)\\s*;"
	cp := regexp.MustCompile(package_pattern)
	for i, line := range p.content {
		if parser.IsExtendType(line) {
			return
		}
		matches := cp.FindStringSubmatch(line)
		if len(matches) > 1 {
			p.Package = matches[1]
			p.Start += i
		}
	}
}

func (p *Proto) initImports() {
	pattern := "^import\\s+\"(.+)\"\\s*;"
	cp := regexp.MustCompile(pattern)
	i := p.Start
	for i < len(p.content) {
		line := p.content[i]
		if parser.IsExtendType(line) {
			return
		}
		matches := cp.FindStringSubmatch(line)
		i++
		if len(matches) > 1 {
			p.Imports = append(p.Imports, matches[1])
			p.Start += 1
			continue
		}
	}
}

func (p *Proto) IsSupported() bool {
	return p.syntax() == PROTO_SYNTAX_3
}

func (p *Proto) Parse() {
	total := len(p.content)
	//var messages []*Message
	var i int = p.Start
	for {
		if i >= total {
			break
		}

		line := p.content[i]

		comment, step := parser.ReadComment(p.content[i:])
		if step > 0 {
			i += step
			line = p.content[i]
		}
		//import, to move cursor to below of message
		i++

		if parser.StartWithMessage(line) {
			name := parser.GetMessageName(line)
			message := &Message{
				Name:    name,
				Comment: comment,
				Package: p.Package,
			}
			//skip := ParseMessage(p.content[i:], 1, message)
			skip := message.Parse(p.content[i:], 1)
			//message.WriteHtml()
			i += skip
			p.Messages = append(p.Messages, message)
		} else if parser.StartWithService(line) {
			service := &Service{
				Name:    parser.GetServiceName(line),
				Note:    comment,
				Package: p.Package,
			}
			skip := service.Parse(p.content[i:], 1)
			i += skip
			//service.WriteHtml()
			p.Services = append(p.Services, service)
		}
	}
}

func (p *Proto) JSON() (string, error) {
	buf, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return "", err
	}
	return string(buf), err
}

func (p *Proto) ResolveImports() []*Proto {
	var protos []*Proto
	for _, importFile := range p.Imports {
		ps := p.ResolveImport(importFile)
		if ps != nil {
			protos = append(protos, ps...)
		}
	}
	return protos
}

func (p *Proto) ResolveImport(importFile string) []*Proto {
	protoFile := GetAbsPath(p.Package, importFile)
	if IsProtoParsed(p.Package, importFile) {
		return nil
	}
	RegisterProto(p.Package, importFile)

	newProto := &Proto{}
	return newProto.Initialize(protoFile)
}
