package pb

import (
	"github.com/wangming1993/pb2doc/parser"
)

type Service struct {
	Package string
	Name    string
	RPCs    []*RPC
	Note    string
}

func (s *Service) Parse(lines []string, depth int) int {
	total := len(lines)
	i := 0

	for {

		if i >= total {
			break
		}
		line := lines[i]

		if parser.EndWithBrace(line) {
			depth--
			if depth == 0 {
				return i
			}
		}

		comment, fs := parser.ReadComment(lines[i:])
		if fs > 0 {
			i += fs
			line = lines[i]
		}
		i++

		rpc := NewRPCWithNote(s.Package, line, comment)
		if rpc != nil {
			s.RPCs = append(s.RPCs, rpc)
		}

	}
	return i
}
