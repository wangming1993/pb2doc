package pb

import (
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/wangming1993/pb2doc/parser"
)

var messageTemplate string = "templates/message.mustache"
var serviceTemplate string = "templates/service.mustache"

func (m *Message) WriteHtml() error {
	err := m.html()
	for _, ms := range m.Messages {
		err = ms.html()
	}
	return err
}

func (m *Message) html() error {
	for _, f := range m.Fields {
		f.WithLink("")
	}
	out, _ := mustache.RenderFile(messageTemplate,
		map[string]interface{}{
			"Name":    m.Name,
			"Comment": parser.PrettifyNote(m.Comment),
			"Fields":  m.Fields,
		},
	)

	pkgs := append([]string{"htmls"}, strings.Split(m.Package, ".")...)
	path := filepath.Join(pkgs...)
	name := m.Name + ".html"
	file, err := parser.CreateFile(path, name)
	if err != nil {
		return err
	}
	_, err = file.WriteString(out)
	return err
}

func (s *Service) WriteHtml() error {
	/*for _, rpc := range s.RPCs {
		//f.WithLink("")
	}*/
	out, _ := mustache.RenderFile(serviceTemplate,
		map[string]interface{}{
			"Name": "Service",
			"Note": parser.PrettifyNote(s.Comment),
			"RPCs": s.RPCs,
		},
	)

	pkgs := append([]string{"htmls"}, strings.Split(s.Package, ".")...)
	path := filepath.Join(pkgs...)
	name := "services.html"
	file, err := parser.CreateFile(path, name)
	if err != nil {
		return err
	}
	_, err = file.WriteString(out)
	return err
}
