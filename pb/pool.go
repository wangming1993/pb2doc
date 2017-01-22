package pb

import (
	"path"
	"strings"

	"github.com/wangming1993/pb2doc/parser"
)

var ParsedMessages map[string]*Message
var ParsingProto map[string]bool

func init() {
	ParsedMessages = make(map[string]*Message)
	ParsingProto = make(map[string]bool)
}

func MakeParsed(pkg, name string, msg *Message) {
	abs := pkg + "." + name
	ParsedMessages[abs] = msg
}

func IsProtoParsed(pkg, name string) bool {
	abs := GetAbsPath(pkg, name)
	if parsed, ok := ParsingProto[abs]; ok {
		return parsed
	}
	return false
}

func RegisterProto(pkg, name string) {
	abs := GetAbsPath(pkg, name)
	ParsingProto[abs] = true

}

// GetAbsPackage used to get absolute package name of one proto file
// pkg means package name
// imp means import name, one import has such format:
//                 import "common/types/types.proto"
func GetAbsPackage(pkg, imp string) string {
	ps := strings.Split(imp, ".")
	total := len(ps)
	if total > 1 {
		return imp
	}
	return pkg + "." + ps[total-1]
}

func GetAbsPath(pkg, imp string) string {
	var rel string
	ps := strings.Split(imp, "/")
	total := len(ps)
	if total > 1 {
		rel = imp
	} else {
		rel = strings.Replace(pkg, ".", "/", -1) + "/" + ps[total-1]
	}
	rel = strings.TrimPrefix(rel, parser.GetPKGPrefix()+"/")
	return path.Join(parser.GetBasePath(), rel)
}
