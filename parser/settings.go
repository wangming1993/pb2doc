package parser

var ProtoPath string
var PackagePrefix string

func SetBasePath(path string) {
	ProtoPath = path
}

func GetBasePath() string {
	return ProtoPath
}

func SetPrefix(prefix string) {
	PackagePrefix = prefix
}

func GetPKGPrefix() string {
	return PackagePrefix
}
