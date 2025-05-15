package expr

import "fmt"

type Type string

const (
	TypeString Type = "string"
	TypeInt    Type = "int"
	TypeUint   Type = "uint"
	TypeFloat  Type = "float"
	TypeBool   Type = "bool"
)

func (t *Type) UnmarshalText(b []byte) error {
	switch string(b) {
	case string(TypeString):
		*t = TypeString
	case string(TypeInt):
		*t = TypeInt
	case string(TypeUint):
		*t = TypeUint
	case string(TypeFloat):
		*t = TypeFloat
	case string(TypeBool):
		*t = TypeBool
	default:
		return fmt.Errorf("invalid type '%s'", string(b))
	}
	return nil
}
