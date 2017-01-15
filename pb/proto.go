package pb

import (
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
	start    int // the message start int, excludes syntax,package and imports line
	Syntax   string
}

func (p *Proto) Initialize(file string) error {
	fmt.Println("initialize...")
	lines := parser.ReadFile(file)
	p.content = lines
	supported := p.IsSupported()
	if !supported {
		return errors.New("Unsupported proto syntax...")
	}
	p.Syntax = p.syntax()
	p.start = 1 //excludes syntax line

	p.initPackage()
	p.ParseMessage()
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
		matches := cp.FindStringSubmatch(line)
		if len(matches) > 1 {
			p.Package = matches[1]
			p.start += i
			fmt.Println(p.Package, p.start)
		}
	}
}

func (p *Proto) IsSupported() bool {
	return p.syntax() == PROTO_SYNTAX_3
}

func (p *Proto) ParseMessage() {
	total := len(p.content)
	//var messages []*Message
	var i int = p.start
	for {
		if i >= total {
			break
		}

		line := p.content[i]

		comment, step := parser.ReadComment(p.content[i:])
		if step > 0 {
			i += step
			line = p.content[i]
		} else {
			i++
		}

		if parser.StartWithMessage(line) {
			name := parser.GetMessageName(line)
			message := &Message{
				Name:    name,
				Comment: comment,
				Package: p.Package,
			}

			for {
				if parser.EndWithBrace(line) || i >= total {
					break
				}

				fc, fs := parser.ReadComment(p.content[i:])
				line = p.content[i]

				if fs > 0 {
					i += fs
					line = p.content[i]
				}
				field := NewFieldWithComment(line, fc)
				if field != nil {
					message.Fields = append(message.Fields, field)
				}
				i++
			}
			content, _ := message.JSON()
			fmt.Println(content)
		}
	}
}
