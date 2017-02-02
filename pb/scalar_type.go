package pb

const (
	PROTO_SYNTAX_3       = "proto3"
	PROTO_SYNTAX_UNKNOWN = "unknown"
)

type ScalarType string

const (
	DOUBLE ScalarType = "double"
	FLOAT  ScalarType = "float"
	INT32  ScalarType = "int32"
	INT64  ScalarType = "int64"
	BOOL   ScalarType = "bool"
	STRING ScalarType = "string"
	MAP    ScalarType = "map"
)

var scalarTypes map[string]ScalarType = map[string]ScalarType{
	"double": DOUBLE,
	"float":  FLOAT,
	"int32":  INT32,
	"int64":  INT64,
	"bool":   BOOL,
	"string": STRING,
	"map":    MAP,
}

func IsScalarType(name string) bool {
	if _, ok := scalarTypes[name]; ok {
		return true
	}
	return false
}
