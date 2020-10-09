package event

type MessageBus struct {
	Driver EventDriver
	messageDriver
}

func (b *MessageBus) SetDefaults() {
	if b.Driver == EVENT_DRIVER_UNKNOWN {
		b.Driver = EVENT_DRIVER__BUILDIN
	}
}

func (b *MessageBus) Init() {
	if b.messageDriver != nil {
		return
	}
	switch b.Driver {
	case EVENT_DRIVER__BUILDIN:
		b.messageDriver = newMemoryMessageBus()
	default:
		panic("[MessageBus] Driver must be defined")
	}
}
