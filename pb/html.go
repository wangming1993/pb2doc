package pb

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/wangming1993/pb2doc/parser"
)

var indexTemplate string = "templates/index.mustache"
var serviceTemplate string = "templates/service.mustache"
var messageServiceTemplate string = "templates/message_service.mustache"
var enumServiceTemplate string = "templates/enum_service.mustache"

func (m *Message) WriteHtmlWithNavigator(basePath string, navigators []*Navigator) error {
	for _, f := range m.Fields {
		f.WithLink("")
	}

	path := getPath(basePath, m.Package)
	return writeHTMLFile(path, m.Name, messageServiceTemplate, map[string]interface{}{
		"Name":       m.Name,
		"Note":       parser.PrettifyNote(m.Note),
		"Fields":     m.Fields,
		"Navigators": navigators,
	})
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

	path := getPath(basePath, s.Package)
	return writeHTMLFile(path, s.Name, serviceTemplate, map[string]interface{}{
		"Name": s.Name,
		"Note": parser.PrettifyNote(s.Note),
		"RPCs": s.RPCs,
	})
}

func (s *Service) Position() string {
	pkg := strings.TrimPrefix(s.Package, parser.GetPKGPrefix())
	pkg = strings.TrimPrefix(pkg, ".")
	pkgs := strings.Split(pkg, ".")
	path := filepath.Join(pkgs...)
	name := s.Name + ".html"
	return filepath.Join(path, name)
}

func (e *Enum) WriteHtmlWithNavigator(basePath string, navigators []*Navigator) error {
	path := getPath(basePath, e.pkg)
	return writeHTMLFile(path, e.Name, enumServiceTemplate, map[string]interface{}{
		"Name":       e.Name,
		"Note":       parser.PrettifyNote(e.Note),
		"Elems":      e.Elems,
		"Navigators": navigators,
	})
}

func (e *Enum) WriteHtmlWithService(basePath string, services []*Service) error {
	var navs []*Navigator
	for _, s := range services {
		navs = append(navs, NewNavigator(s, e.pkg))
	}
	return e.WriteHtmlWithNavigator(basePath, navs)
}

func GenerateIndexHTML(basePath string, services []*Service) error {
	pkgPrefix := parser.GetPKGPrefix()
	path := strings.Join([]string{basePath, pkgPrefix}, "/")
	var renderedServices []map[string]string
	for _, service := range services {
		renderedServices = append(renderedServices, map[string]string{
			"Name":   service.Name,
			"Folder": strings.Replace(service.Package, pkgPrefix+".", "", -1),
			"Count":  strconv.Itoa(len(service.RPCs)),
		})
	}
	return writeHTMLFile(path, "index", indexTemplate, map[string]interface{}{
		"Services": renderedServices,
	})
}

func writeHTMLFile(path, fileName, tmplName string, data map[string]interface{}) error {
	file, err := parser.CreateFile(path, fileName+".html")
	if err != nil {
		return err
	}
	content, _ := mustache.RenderFile(tmplName, data)
	_, err = file.WriteString(content)
	return err
}

func getPath(basePath, namespace string) string {
	pkgs := append([]string{basePath}, strings.Split(namespace, ".")...)
	return filepath.Join(pkgs...)
}
