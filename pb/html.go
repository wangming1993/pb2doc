package pb

import (
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/wangming1993/pb2doc/parser"
)

var messageTemplate string = "templates/message.mustache"
var serviceTemplate string = "templates/service.mustache"
var messageServiceTemplate string = "templates/message_service.mustache"
var enumServiceTemplate string = "templates/enum_service.mustache"

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

func (m *Message) WriteHtmlWithNavigator(basePath string, navigators []*Navigator) error {
	for _, f := range m.Fields {
		f.WithLink("")
	}

	out, _ := mustache.RenderFile(messageServiceTemplate,
		map[string]interface{}{
			"Name":       m.Name,
			"Comment":    parser.PrettifyNote(m.Comment),
			"Fields":     m.Fields,
			"Navigators": navigators,
		},
	)

	err := writeHtmlFile(basePath, m.Package, m.Name, out)
	if err != nil {
		return err
	}
	return nil
}

func (m *Message) WriteHtmlWithService(basePath string, services []*Service) error {
	var navs []*Navigator
	for _, s := range services {
		navs = append(navs, NewNavigator(s, m.Package))
	}
	return m.WriteHtmlWithNavigator(basePath, navs)
}

func (s *Service) WriteHtml(basePath string) error {
	for _, rpc := range s.RPCs {
		rpc.WithLink("")
	}
	out, _ := mustache.RenderFile(serviceTemplate,
		map[string]interface{}{
			"Name": s.Name,
			"Note": parser.PrettifyNote(s.Note),
			"RPCs": s.RPCs,
		},
	)

	err := writeHtmlFile(basePath, s.Package, s.Name, out)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Position() string {
	pkg := strings.TrimPrefix(s.Package, parser.GetPKGPrefix())
	pkg = strings.TrimPrefix(pkg, ".")
	pkgs := strings.Split(pkg, ".")
	path := filepath.Join(pkgs...)
	name := s.Name + ".html"
	return filepath.Join(path, name)
}

func writeHtmlFile(basePath, namespace, name, out string) error {
	pkgs := append([]string{basePath}, strings.Split(namespace, ".")...)
	path := filepath.Join(pkgs...)
	fileName := name + ".html"
	file, err := parser.CreateFile(path, fileName)
	if err != nil {
		return err
	}
	_, err = file.WriteString(out)
	return err
}

func (e *Enum) WriteHtmlWithNavigator(basePath string, navigators []*Navigator) error {
	out, _ := mustache.RenderFile(enumServiceTemplate,
		map[string]interface{}{
			"Name":       e.Name,
			"Comment":    parser.PrettifyNote(e.Note),
			"Elems":      e.Elems,
			"Navigators": navigators,
		},
	)

	err := writeHtmlFile(basePath, e.pkg, e.Name, out)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enum) WriteHtmlWithService(basePath string, services []*Service) error {
	var navs []*Navigator
	for _, s := range services {
		navs = append(navs, NewNavigator(s, e.pkg))
	}
	return e.WriteHtmlWithNavigator(basePath, navs)
}
