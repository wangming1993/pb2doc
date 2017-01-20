package pb

import (
	"os"

	"github.com/cbroglie/mustache"
	"github.com/wangming1993/pb2doc/parser"
)

var template string = "templates/message.mustache"

func (m *Message) WriteHtml() error {
	out, _ := mustache.RenderFile(template,
		map[string]interface{}{
			"Name":    m.Name,
			"Comment": parser.PrettifyNote(m.Comment),
			"Fields":  m.Fields,
		},
	)

	file, err := os.Create("htmls/" + m.Name + ".html")
	if err != nil {
		return err
	}

	_, err = file.WriteString(out)
	return err
}
