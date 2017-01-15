package pb

import (
	"fmt"
	"github.com/spf13/cast"
	"regexp"
)

type Field struct {
	Type    string
	Name    string
	Order   int
	Comment string
}

func NewField(line string) *Field {
	pattern := "^\\s+([a-z0-9]+)\\s+([a-z0-9_-]+)\\s?=\\s?([0-9]+)\\s*;\\s*((//.*)|(/\\*.*\\*/))?"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 4 {
		return nil
	}
	field := &Field{
		Type:  matches[1],
		Name:  matches[2],
		Order: cast.ToInt(matches[3]),
	}
	if len(matches) > 4 {
		field.Comment = matches[4]
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
