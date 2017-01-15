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
}

func (p *Proto) Initialize(file string) error {
	fmt.Println("initialize...")
	lines := parser.ReadFile(file)
	p.content = lines
	supported := p.IsSupported()
	if !supported {
		return errors.New("Unsupported proto syntax...")
	}
	p.ParseMessage()
	return nil
}

func (p *Proto) syntax() string {
	syntax_pattern := "syntax\\s?=\\s?\"(.+)\"\\s?;"
	cp := regexp.MustCompile(syntax_pattern)
	matches := cp.FindStringSubmatch(p.content[0])
	if len(matches) > 1 {
		return matches[1]
	}
	return PROTO_SYNTAX_UNKNOWN
}

func (p *Proto) IsSupported() bool {
	return p.syntax() == PROTO_SYNTAX_3
}

func (p *Proto) ParseMessage() {
	total := len(p.content)
	//var messages []*Message
	var i int = 0
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
