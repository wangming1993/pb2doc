package pb

import (
	"github.com/wangming1993/pb2doc/parser"
)

type Oneof struct {
	pkg    string
	Name   string
	Note   string
	Fields []*Field
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

		note, fs := parser.ReadNote(lines[i:])
		if fs > 0 {
			i += fs
			line = lines[i]
		}

		field := NewFieldWithNote(oneof.pkg, line, note)
		if field != nil {
			oneof.Fields = append(oneof.Fields, field)
		}
		i++

	}
	return i
}
