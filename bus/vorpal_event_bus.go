package bus

type vorpalEventBus struct {
	keyEventListenerChannels             []chan KeyEvent
	mouseListenerChannels                []chan MouseEvent
	drawEventListenerChannels            []chan DrawEvent
	audioEventListenerChannels           []chan AudioEvent
	textEventListenerChannels            []chan TextEvent
	keysRegistrationEventListenerChannel []chan KeysRegistrationEvent
	controlEventListenerChannel          []chan ControlEvent
}

type VorpalBus interface {

	///Individual listeners....
	AddKeyEventListener(eventListener KeyEventListener)
	AddMouseListener(eventListener MouseEventListener)
	AddDrawEventListener(eventListener DrawEventListener)
	AddAudioEventListener(eventListener AudioEventListener)
	AddTextEventListener(eventListener TextEventListener)
	AddKeysRegistrationEventListener(eventListener KeysRegistrationEventListener)
	AddControlEventListener(eventListener ControlEventListener)

	SendMouseEvent(event MouseEvent)
	SendKeysRegistrationEvent(event KeysRegistrationEvent)
	SendKeyEvent(event KeyEvent)
	SendDrawEvent(event DrawEvent)
	SendAudioEvent(event AudioEvent)
	SendTextEvent(event TextEvent)
	SendControlEvent(event ControlEvent)
}

var eb = vorpalEventBus{}

func GetVorpalBus() VorpalBus {
	return &eb
}

// Channels that can buffer multiple events or where we don't care for only the latest
// event. For example, if the 10 mouse events are sent but the consumer is only concerned
// with the last one, they can ignore all but the last. We don't want to block the caller.
//
// These also give the consumer the choice whether to igonre or process all the events. If there are
// multiple DrawEvents we may wish to process them al in one context while ignore all but the latest
// and consider the earlier  ones to be frame misses.
// The controller determines correct behavior
// Practical limit of 10 though as more indicates a lot of processing misses.

func (bus *vorpalEventBus) AddControlEventListener(eventListener ControlEventListener) {
	listenerChannel := make(chan ControlEvent, 10)
	bus.controlEventListenerChannel = append(bus.controlEventListenerChannel, listenerChannel)
	go eventListener.OnControlEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendControlEvent(event ControlEvent) {
	for _, channel := range bus.controlEventListenerChannel {
		channel <- event
	}
}
func (bus *vorpalEventBus) AddKeysRegistrationEventListener(eventListener KeysRegistrationEventListener) {
	listenerChannel := make(chan KeysRegistrationEvent, 10)
	bus.keysRegistrationEventListenerChannel = append(bus.keysRegistrationEventListenerChannel, listenerChannel)
	go eventListener.OnKeyRegistrationEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendKeysRegistrationEvent(event KeysRegistrationEvent) {
	for _, channel := range bus.keysRegistrationEventListenerChannel {
		channel <- event
	}
}

func (bus *vorpalEventBus) AddTextEventListener(eventListener TextEventListener) {
	listenerChannel := make(chan TextEvent, 10)
	bus.textEventListenerChannels = append(bus.textEventListenerChannels, listenerChannel)
	go eventListener.OnTextEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendTextEvent(event TextEvent) {
	for _, channel := range bus.textEventListenerChannels {
		channel <- event
	}
}

func (bus *vorpalEventBus) AddAudioEventListener(eventListener AudioEventListener) {
	listenerChannel := make(chan AudioEvent, 10)
	bus.audioEventListenerChannels = append(bus.audioEventListenerChannels, listenerChannel)
	go eventListener.OnAudioEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendAudioEvent(event AudioEvent) {
	for _, channel := range bus.audioEventListenerChannels {
		channel <- event
	}
}

// /KEY EVENTS
func (bus *vorpalEventBus) AddKeyEventListener(eventListener KeyEventListener) {
	listenerChannel := make(chan KeyEvent, 10)
	bus.keyEventListenerChannels = append(bus.keyEventListenerChannels, listenerChannel)
	go eventListener.OnKeyEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendKeyEvent(event KeyEvent) {
	for _, channel := range bus.keyEventListenerChannels {
		channel <- event
	}
}

// ///MOUSE BUTTON EVENTS
func (bus *vorpalEventBus) AddMouseListener(eventListener MouseEventListener) {
	listenerChannel := make(chan MouseEvent, 10)
	bus.mouseListenerChannels = append(bus.mouseListenerChannels, listenerChannel)
	go eventListener.OnMouseEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendMouseEvent(event MouseEvent) {
	for _, channel := range bus.mouseListenerChannels {
		channel <- event
	}
}

func (bus *vorpalEventBus) AddDrawEventListener(eventListener DrawEventListener) {
	listenerChannel := make(chan DrawEvent, 10)
	bus.drawEventListenerChannels = append(bus.drawEventListenerChannels, listenerChannel)
	go eventListener.OnDrawEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendDrawEvent(event DrawEvent) {
	for _, channel := range bus.drawEventListenerChannels {
		channel <- event
	}
}

//
