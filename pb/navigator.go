package pb

import (
	"path/filepath"
	"strings"
)

type Navigator struct {
	Position string
	Name     string
}

func NewNavigator(service *Service, message *Message) *Navigator {
	nav := &Navigator{
		Name: service.Name,
	}
	position := service.Position()
	pkg := message.Package

	var ps []string

	pkgs := strings.Split(pkg, ".")
	depth := len(pkgs)
	if depth > 1 {
		if pkg != "" {
			for i := 0; i < depth; i++ {
				ps = append(ps, "..")
			}
		}
	}
	ps = append(ps, position)
	nav.Position = filepath.Join(ps...)
	return nav
}
