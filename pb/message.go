package pb

import (
	"github.com/spf13/cast"
	"github.com/wangming1993/pb2doc/parser"
)

type Message struct {
	Note     string
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

func (m *Message) GetAll() []*Message {
	var messages []*Message = []*Message{m}
	for _, message := range m.Messages {
		messages = append(messages, message.GetAll()...)
	}
	return messages
}

func (message *Message) Parse(lines []string, depth int) int {
	total := len(lines)
	i := 0

	for {

		if i >= total {
			break
		}
		line := lines[i]

		if parser.EndWithBrace(line) {
			//log.Println(line)
			depth--
			if depth == 0 {
				return i
			}
		}

		note, fs := parser.ReadNote(lines[i:])
		if fs > 0 {
			i += fs
			line = lines[i]
		}
		i++

		if parser.IsExtendType(line) {
			depth++
			if parser.StartWithMessage(line) {
				embedMessage := &Message{
					Name:    parser.GetMessageName(line),
					Note:    note,
					Package: message.Package,
				}
				i += embedMessage.Parse(lines[i:], 1)
				message.Messages = append(message.Messages, embedMessage)
			} else if parser.StartWithEnum(line) {
				embedEnum := &Enum{
					Name: parser.GetEnumName(line),
					Note: note,
					pkg:  message.Package,
				}
				embedEnum.Parse(lines)
				message.Enums = append(message.Enums, embedEnum)
			} else if parser.StartWithOneof(line) {
				embedOneof := &Oneof{
					Name: parser.GetOneofName(line),
					Note: note,
				}
				message.Oneofs = append(message.Oneofs, embedOneof)

				step := ParseOneof(lines[i:], embedOneof)
				i += step
				message.Fields = append(message.Fields, embedOneof.Fields...)
			}
		} else {
			var field *Field
			if parser.StartWithMap(line) {
				field = NewMapField(message.Package, line, note)
			} else {
				field = NewFieldWithNote(message.Package, line, note)
			}
			if field != nil {
				message.Fields = append(message.Fields, field)
			}
		}
	}
	return i
}
