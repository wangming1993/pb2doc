package pb

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/wangming1993/pb2doc/parser"
	"regexp"
)

type Field struct {
	Modifier string
	Type     string
	Name     string
	Order    int
	Comment  string
}

func NewField(line string) *Field {
	pattern := "^\\s*([a-z]*)\\s+([a-z0-9]+)\\s+([a-z0-9_-]+)\\s?=\\s?([0-9]+)\\s*;\\s*((//.*)|(/\\*.*\\*/))?"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 5 {
		return nil
	}
	field := &Field{
		Modifier: matches[1],
		Type:     matches[2],
		Name:     matches[3],
		Order:    cast.ToInt(matches[4]),
	}
	if len(matches) > 5 {
		field.Comment = parser.PrettifyNote(matches[5])
	}
	return field
}

func NewFieldWithComment(line, comment string) *Field {
	field := NewField(line)
	if field == nil {
		return nil
	}
	if comment != "" {
		fmt.Println(comment)
		fmt.Println(parser.PrettifyNote(comment))

		field.Comment = parser.PrettifyNote(comment)
	}
	return field
}

func (f *Field) String() string {
	return fmt.Sprintf("Field:%s, type:%s, order:%d, comment:%s", f.Name, f.Type, f.Order, f.Comment)
}
