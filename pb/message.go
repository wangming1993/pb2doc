package pb

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type Message struct {
	Fields  []*Field
	Comment string
	Package string
	Name    string
}

func (m *Message) String() string {
	return cast.ToString(len(m.Fields))
}

func (m *Message) Data() {
	fmt.Println("--------------------------------------")
	fmt.Println("Name is:\n", m.Name)
	fmt.Println("Comments is:\n", m.Comment)
	fmt.Println("Fields count:\n", len(m.Fields))
	for _, f := range m.Fields {
		f.String()
	}
	fmt.Println("--------------------------------------")
}

func (m *Message) JSON() (string, error) {
	buf, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return "", err
	}
	return string(buf), err
}
