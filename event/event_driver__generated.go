package event

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidEventDriver = errors.New("invalid EventDriver")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("EventDriver", map[string]string{
		"BUILDIN": "内存",
	})
}

func ParseEventDriverFromString(s string) (EventDriver, error) {
	switch s {
	case "":
		return EVENT_DRIVER_UNKNOWN, nil
	case "BUILDIN":
		return EVENT_DRIVER__BUILDIN, nil
	}
	return EVENT_DRIVER_UNKNOWN, InvalidEventDriver
}

func ParseEventDriverFromLabelString(s string) (EventDriver, error) {
	switch s {
	case "":
		return EVENT_DRIVER_UNKNOWN, nil
	case "内存":
		return EVENT_DRIVER__BUILDIN, nil
	}
	return EVENT_DRIVER_UNKNOWN, InvalidEventDriver
}

func (EventDriver) EnumType() string {
	return "EventDriver"
}

func (EventDriver) Enums() map[int][]string {
	return map[int][]string{
		int(EVENT_DRIVER__BUILDIN): {"BUILDIN", "内存"},
	}
}

func (v EventDriver) String() string {
	switch v {
	case EVENT_DRIVER_UNKNOWN:
		return ""
	case EVENT_DRIVER__BUILDIN:
		return "BUILDIN"
	}
	return "UNKNOWN"
}

func (v EventDriver) Label() string {
	switch v {
	case EVENT_DRIVER_UNKNOWN:
		return ""
	case EVENT_DRIVER__BUILDIN:
		return "内存"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*EventDriver)(nil)

func (v EventDriver) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidEventDriver
	}
	return []byte(str), nil
}

func (v *EventDriver) UnmarshalText(data []byte) (err error) {
	*v, err = ParseEventDriverFromString(string(bytes.ToUpper(data)))
	return
}
