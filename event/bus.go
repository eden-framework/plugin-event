package event

type MessageBus struct {
	Driver DriverType
	messageDriver
}

func (b *MessageBus) SetDefaults() {
	if b.Driver == DRIVER_TYPE_UNKNOWN {
		b.Driver = DRIVER_TYPE__MEMORY
	}
}

func (b *MessageBus) Init() {
	if b.messageDriver != nil {
		return
	}
	switch b.Driver {
	case DRIVER_TYPE__MEMORY:
		b.messageDriver = newMemoryMessageBus()
	default:
		panic("[MessageBus] Driver must be defined")
	}
}
