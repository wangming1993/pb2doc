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
)
