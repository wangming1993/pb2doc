package pb

import (
	"path/filepath"
	"strings"

	"github.com/wangming1993/pb2doc/parser"
)

type Navigator struct {
	Position string
	Name     string
}

func NewNavigator(service *Service, pkg string) *Navigator {
	pkg = parser.TrimPKGPrefix(pkg)
	nav := &Navigator{
		Name: service.Name,
	}
	position := service.Position()

	var ps []string

	pkgs := strings.Split(pkg, ".")
	depth := len(pkgs)
	if pkg != "" {
		for i := 0; i < depth; i++ {
			ps = append(ps, "..")
		}
	}
	ps = append(ps, position)
	nav.Position = filepath.Join(ps...)
	return nav
}
