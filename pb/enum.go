package pb

import (
	"fmt"
	"regexp"

	"github.com/spf13/cast"
	"github.com/wangming1993/pb2doc/parser"
)

type Enum struct {
	Note  string
	Name  string
	Elems []*Elem
	pkg   string
}

func (e *Enum) Data() {
	fmt.Println("Name:", e.Name)
	fmt.Println("    elem:")

	for _, elem := range e.Elems {
		fmt.Printf("     %s = %d \n", elem.Value, elem.Order)
	}
}

type Elem struct {
	Value string
	Order int
	Note  string
}

func NewElemWithNote(line, note string) *Elem {
	pattern := "^\\s*([A-Za-z0-9_-]+)\\s?=\\s?([0-9]+)\\s*;\\s*((//.*)|(/\\*.*\\*/))?"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 3 {
		return nil
	}

	elem := &Elem{
		Value: matches[1],
		Order: cast.ToInt(matches[2]),
	}

	if len(matches) > 3 && matches[3] != "" {
		note = matches[3]
	}

	if note != "" {
		elem.Note = parser.PrettifyNote(note)
	}

	return elem
}

func (e *Enum) Parse(lines []string) int {
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
		i++

		elem := NewElemWithNote(line, note)
		if elem != nil {
			e.Elems = append(e.Elems, elem)
		}

	}
	return i
}
