package pb

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cast"
	"github.com/wangming1993/pb2doc/parser"
)

type Field struct {
	Label string
	Type  string
	Name  string
	Order int
	Note  string
	Link  string
	pkg   string
}

func NewMapField(pkg, line, note string) *Field {
	pattern := "^\\s*map\\s*<\\s*([a-zA-Z0-9_.-]+)\\s*,\\s*([a-zA-Z0-9_.-]+)\\s*>\\s*([a-z0-9_-]+)\\s?=\\s?([0-9]+)\\s*;\\s*((//.*)|(/\\*.*\\*/))?"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 5 {
		return nil
	}
	field := &Field{
		Label: "optional",
		Type:  "map",
		Name:  matches[3],
		Order: cast.ToInt(matches[4]),
		pkg:   pkg,
	}
	if len(matches) > 5 {
		field.Note = parser.PrettifyNote(matches[5])
	}

	return field
}

func NewField(pkg, line string) *Field {
	pattern := "^\\s*([a-z]*)\\s+([a-zA-Z0-9.]+)\\s+([a-z0-9_-]+)\\s?=\\s?([0-9]+)\\s*;\\s*((//.*)|(/\\*.*\\*/))?"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 5 {
		return nil
	}
	label := matches[1]
	if label == "" {
		label = "optional"
	}
	field := &Field{
		Label: label,
		Type:  matches[2],
		Name:  matches[3],
		Order: cast.ToInt(matches[4]),
		pkg:   pkg,
	}
	if len(matches) > 5 {
		field.Note = parser.PrettifyNote(matches[5])
	}

	return field
}

func NewFieldWithNote(pkg, line, note string) *Field {
	field := NewField(pkg, line)
	if field == nil {
		return nil
	}
	if note != "" {
		field.Note = parser.PrettifyNote(note)
	}
	return field
}

func (f *Field) String() string {
	return fmt.Sprintf("Field:%s, type:%s, order:%d, note:%s", f.Name, f.Type, f.Order, f.Note)
}

func (f *Field) WithLink(path string) {
	if IsScalarType(f.Type) {
		return
	}
	pkg := f.formatDot(f.pkg)
	ps := []string{path}
	ts := strings.Split(f.Type, ".")
	length := len(ts)
	fileName := ts[length-1] + ".html"
	if length > 1 {
		if pkg != "" {
			pkgs := strings.Split(pkg, ".")
			for i := 0; i <= len(pkgs); i++ {
				ps = append(ps, "..")
			}
		}
		ps = append(ps, ts[0:length-1]...)
	}
	ps = append(ps, fileName)
	f.Link = filepath.Join(ps...)
	f.Type = f.formatDot(f.Type)
}

func (f *Field) formatDot(withDot string) string {
	fieldType := strings.TrimPrefix(withDot, parser.GetPKGPrefix())
	return strings.TrimPrefix(fieldType, ".")
}
