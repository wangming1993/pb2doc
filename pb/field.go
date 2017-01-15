package pb

import (
	"fmt"
	"github.com/spf13/cast"
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
		field.Comment = matches[5]
	}
	return field
}

func NewFieldWithComment(line, comment string) *Field {
	field := NewField(line)
	if field == nil {
		return nil
	}
	if comment != "" {
		field.Comment = comment
	}
	return field
}

func (f *Field) String() {
	fmt.Printf("Field:%s, type:%s, order:%d, comment:%s \n", f.Name, f.Type, f.Order, f.Comment)
}
