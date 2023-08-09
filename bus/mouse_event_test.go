package bus

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/mobile/event/mouse"
	"testing"
)

//TODO Think about the design and use generics.

func TestMouseBasics(t *testing.T) {
	evt := createBasicEvent()
	mouseEvent := NewMouseEvent(evt)
	assert.Equal(t, &cursorPoint{100, 200}, mouseEvent.GetCursorPoint())
	assert.False(t, mouseEvent.IsPressed(mouse.ButtonMiddle))
	assert.False(t, mouseEvent.IsPressed(mouse.ButtonRight))
	assert.False(t, mouseEvent.IsPressed(mouse.ButtonLeft))
	evt.Button = mouse.ButtonLeft
	evt.Direction = mouse.DirPress
	//NewEvent since we are passing by value...
	mouseEvent = NewMouseEvent(evt)
	assert.True(t, mouseEvent.IsPressed(mouse.ButtonLeft))
}

func createBasicEvent() mouse.Event {
	return mouse.Event{X: 100, Y: 200, Button: 0, Modifiers: 0, Direction: 0}
}
