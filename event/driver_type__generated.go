package event

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDriverType = errors.New("invalid DriverType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("DriverType", map[string]string{
		"BUILDIN": "内存",
	})
}

func ParseDriverTypeFromString(s string) (DriverType, error) {
	switch s {
	case "":
		return DRIVER_TYPE_UNKNOWN, nil
	case "BUILDIN":
		return DRIVER_TYPE__BUILDIN, nil
	}
	return DRIVER_TYPE_UNKNOWN, InvalidDriverType
}

func ParseDriverTypeFromLabelString(s string) (DriverType, error) {
	switch s {
	case "":
		return DRIVER_TYPE_UNKNOWN, nil
	case "内存":
		return DRIVER_TYPE__BUILDIN, nil
	}
	return DRIVER_TYPE_UNKNOWN, InvalidDriverType
}

func (DriverType) EnumType() string {
	return "DriverType"
}

func (DriverType) Enums() map[int][]string {
	return map[int][]string{
		int(DRIVER_TYPE__BUILDIN): {"BUILDIN", "内存"},
	}
}

func (v DriverType) String() string {
	switch v {
	case DRIVER_TYPE_UNKNOWN:
		return ""
	case DRIVER_TYPE__BUILDIN:
		return "BUILDIN"
	}
	return "UNKNOWN"
}

func (v DriverType) Label() string {
	switch v {
	case DRIVER_TYPE_UNKNOWN:
		return ""
	case DRIVER_TYPE__BUILDIN:
		return "内存"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*DriverType)(nil)

func (v DriverType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDriverType
	}
	return []byte(str), nil
}

func (v *DriverType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDriverTypeFromString(string(bytes.ToUpper(data)))
	return
}
