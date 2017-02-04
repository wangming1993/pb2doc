package pb

const (
	PROTO_SYNTAX_3       = "proto3"
	PROTO_SYNTAX_UNKNOWN = "unknown"
)

type ScalarType string

const (
	DOUBLE   ScalarType = "double"
	FLOAT    ScalarType = "float"
	INT32    ScalarType = "int32"
	INT64    ScalarType = "int64"
	UINT32   ScalarType = "uint32"
	UINT64   ScalarType = "uint64"
	SINT32   ScalarType = "sint32"
	SINT64   ScalarType = "sint64"
	FIXED32  ScalarType = "fixed32"
	FIXED64  ScalarType = "fixed64"
	SFIXED32 ScalarType = "sfixed32"
	SFIXED64 ScalarType = "sfixed64"
	BOOL     ScalarType = "bool"
	STRING   ScalarType = "string"
	BYTES    ScalarType = "bytes"
	MAP      ScalarType = "map"
)

var scalarTypes map[string]ScalarType = map[string]ScalarType{
	"double":   DOUBLE,
	"float":    FLOAT,
	"int32":    INT32,
	"int64":    INT64,
	"uint32":   UINT32,
	"uint64":   UINT64,
	"sint32":   SINT32,
	"sint64":   SINT64,
	"fixed32":  FIXED32,
	"fixed64":  FIXED64,
	"sfixed32": SFIXED32,
	"sfixed64": SFIXED64,
	"bool":     BOOL,
	"string":   STRING,
	"bytes":    BYTES,
	"map":      MAP,
}

func IsScalarType(name string) bool {
	if _, ok := scalarTypes[name]; ok {
		return true
	}
	return false
}
