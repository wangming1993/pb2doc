package pb

import (
	"fmt"
	"regexp"

	"github.com/wangming1993/pb2doc/parser"
	"path/filepath"
	"strings"
)

type RPC struct {
	Note     string
	Method   string
	Request  string
	Response string
	pkg      string
	ReqLink  string
	RespLink string
}

func NewRPC(pkg, line string) *RPC {
	pattern := "^\\s*rpc\\s+([a-zA-Z0-9_-]+)\\s*\\(([a-z0-9A-Z_\\.-]+)\\)\\s*returns\\s*\\(([0-9a-zA-Z._-]+)\\)\\s*;"
	c, _ := regexp.Compile(pattern)
	matches := c.FindStringSubmatch(line)
	if len(matches) < 4 {
		return nil
	}

	rpc := &RPC{
		Method:   matches[1],
		Request:  matches[2],
		Response: matches[3],
		pkg:      pkg,
	}
	if len(matches) > 4 {
		rpc.Note = parser.PrettifyNote(matches[4])
	}

	return rpc
}

func NewRPCWithNote(pkg, line, note string) *RPC {
	rpc := NewRPC(pkg, line)
	if rpc == nil {
		return nil
	}
	if note != "" {
		rpc.Note = parser.PrettifyNote(note)
	}
	return rpc
}

func (rpc *RPC) String() string {
	return fmt.Sprintf("rpc %s(%s) returns (%s);", rpc.Method, rpc.Request, rpc.Response)
}

func (rpc *RPC) WithLink(path string) {
	rpc.ReqLink = rpc.link(path, rpc.Request)
	rpc.RespLink = rpc.link(path, rpc.Response)

	rpc.Request = rpc.formatDot(rpc.Request)
	rpc.Response = rpc.formatDot(rpc.Response)
}

func (rpc *RPC) link(path, value string) string {
	pkg := rpc.formatDot(rpc.pkg)

	ps := []string{path}
	ts := strings.Split(value, ".")
	length := len(ts)
	fileName := ts[length-1] + ".html"
	if length > 1 {
		if pkg != "" {
			pkgs := strings.Split(pkg, ".")
			for i := 0; i <= len(pkgs); i++ {
				ps = append(ps, "..")
			}
		}
		ps = append(ps, ts[0:length-1]...)
	}
	ps = append(ps, fileName)
	return filepath.Join(ps...)
}

func (rpc *RPC) formatDot(withDot string) string {
	fieldType := strings.TrimPrefix(withDot, parser.GetPKGPrefix())
	return strings.TrimPrefix(fieldType, ".")
}
