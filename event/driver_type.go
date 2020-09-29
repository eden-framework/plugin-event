package event

//go:generate eden generate enum --type-name=DriverType
// api:enum
type DriverType uint8

// driver type
const (
	DRIVER_TYPE_UNKNOWN  DriverType = iota
	DRIVER_TYPE__BUILDIN            // 内存
)
