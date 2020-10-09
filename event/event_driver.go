package event

//go:generate eden generate enum --type-name=EventDriver
// api:enum
type EventDriver uint8

// driver type
const (
	EVENT_DRIVER_UNKNOWN  EventDriver = iota
	EVENT_DRIVER__BUILDIN             // 内存
)
