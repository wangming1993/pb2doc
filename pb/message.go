package pb

import (
	"encoding/json"
	"log"

	"github.com/spf13/cast"
)

type Message struct {
	Comment  string
	Package  string
	Name     string
	Messages []*Message
	Enums    []*Enum
	Oneofs   []*Oneof
	Fields   []*Field
}

func (m *Message) String() string {
	return cast.ToString(len(m.Fields))
}

// Data show the info
func (m *Message) Data() {
	m.WriteHtml()
	log.Println("-----------------------------------------")
	log.Println("Name:", m.Name)
	if len(m.Messages) > 0 {
		log.Println("--> Messages:")
		for _, message := range m.Messages {
			log.Println("------> ", message.Name)
		}
	}
	if len(m.Enums) > 0 {
		log.Println("--> Enums:")
		for _, e := range m.Enums {
			log.Println("------> ", e.Name)
		}
	}
	if len(m.Oneofs) > 0 {
		log.Println("--> Oneofs:")
		for _, e := range m.Oneofs {
			log.Println("------> ", e.Name)
			if len(e.Fields) > 0 {
				log.Println("----------> Fields:")
				for _, f := range e.Fields {
					log.Println("------------> ", f.String())
				}
			}
		}
	}
	if len(m.Fields) > 0 {
		log.Println("--> Fields:")
		for _, f := range m.Fields {
			log.Println("-----> ", f.String())
		}
	}
	log.Println("-----------------------------------------")
}

// JSON use json MarshalIndent to format
func (m *Message) JSON() (string, error) {
	buf, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return "", err
	}
	return string(buf), err
}
