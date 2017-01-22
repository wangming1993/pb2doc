package pb

import (
	"fmt"
	"regexp"

	"github.com/wangming1993/pb2doc/parser"
)

type RPC struct {
	Note     string
	Method   string
	Request  string
	Response string
	pkg      string
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
