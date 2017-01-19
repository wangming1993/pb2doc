package pb

import (
	"github.com/wangming1993/pb2doc/parser"
)

type Oneof struct {
	Name    string
	Comment string
	Fields  []*Field
}

func ParseOneof(lines []string, oneof *Oneof) int {
	total := len(lines)
	i := 0

	for {
		if i >= total {
			break
		}
		line := lines[i]
		if parser.EndWithBrace(line) {
			return i
		}

		comment, fs := parser.ReadComment(lines[i:])
		if fs > 0 {
			i += fs
			line = lines[i]
		}

		field := NewFieldWithComment(line, comment)
		if field != nil {
			oneof.Fields = append(oneof.Fields, field)
		}
		i++

	}
	return i
}
