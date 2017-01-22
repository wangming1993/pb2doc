package pb

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/wangming1993/pb2doc/parser"
)

type Proto struct {
	Messages []*Message
	Comment  string
	Package  string
	content  []string
	Start    int // the message start int, excludes syntax,package and imports line
	Syntax   string
	Imports  []string
	Path     string // base folder of proto file
}

func (p *Proto) Initialize(file string) error {
	lines := parser.ReadFile(file)
	p.content = lines
	supported := p.IsSupported()
	if !supported {
		return errors.New("Unsupported proto syntax...")
	}
	p.Syntax = p.syntax()
	p.Start += 1 //excludes syntax line

	p.initPackage()
	p.initImports()

	fmt.Println(p.Imports, file)

	RegisterProto(p.Package, parser.FileName(file))

	p.ResolveImports()
	p.Parse()

	//fmt.Println(p.JSON())
	return nil
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
			skip := ParseMessage(p.content[i:], 1, message)
			//log.Println(i, skip, i+skip, message.Name)
			message.WriteHtml()
			//message.Data()
			i += skip
		} else if parser.StartWithService(line) {
			service := &Service{
				Name:parser.GetServiceName(line),
			}
			skip := service.Parse(p.content[i:], 1)
			i += skip
			service.WriteHtml()
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

func (p *Proto) ResolveImports() {
	for _, importFile := range p.Imports {
		p.ResolveImport(importFile)
	}
}

func (p *Proto) ResolveImport(importFile string) {
	protoFile := GetAbsPath(p.Package, importFile)
	if IsProtoParsed(p.Package, importFile) {
		return
	}
	RegisterProto(p.Package, importFile)

	newProto := &Proto{}
	newProto.Initialize(protoFile)
}

func ParseMessage(lines []string, depth int, message *Message) int {
	total := len(lines)
	i := 0

	for {

		if i >= total {
			break
		}
		line := lines[i]

		if parser.EndWithBrace(line) {
			//log.Println(line)
			depth--
			if depth == 0 {
				return i
			}
		}

		comment, fs := parser.ReadComment(lines[i:])
		if fs > 0 {
			i += fs
			line = lines[i]
		}
		i++

		if parser.IsExtendType(line) {
			depth++
			//log.Println(line)
			if parser.StartWithMessage(line) {
				embedMessage := &Message{
					Name:    parser.GetMessageName(line),
					Comment: comment,
					Package: message.Package,
				}
				i += ParseMessage(lines[i:], 1, embedMessage)
				message.Messages = append(message.Messages, embedMessage)
			} else if parser.StartWithEnum(line) {
				embedEnum := &Enum{
					Name:    parser.GetEnumName(line),
					Note: comment,
				}
				message.Enums = append(message.Enums, embedEnum)
			} else if parser.StartWithOneof(line) {
				embedOneof := &Oneof{
					Name:    parser.GetOneofName(line),
					Comment: comment,
				}
				message.Oneofs = append(message.Oneofs, embedOneof)

				step := ParseOneof(lines[i:], embedOneof)
				i += step
			}
		} else {
			field := NewFieldWithNote(message.Package, line, comment)
			if field != nil {
				message.Fields = append(message.Fields, field)
			}
		}

	}
	return i
}
