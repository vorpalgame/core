package bus

import (
	"golang.org/x/mobile/event/mouse"
)

func NewMouseEvent(event mouse.Event) MouseEvent {
	return &mouseState{event}
}

type mouseState struct {
	mouse.Event
}
type MouseEventListener interface {
	OnMouseEvent(mouseChannel <-chan MouseEvent)
}

type MouseEvent interface {
	GetCursorPoint() CursorPoint
	IsWheel() bool
	IsReleased(button mouse.Button) bool
	IsPressed(button mouse.Button) bool
}

func (m *mouseState) IsWheel() bool {
	return m.Button.IsWheel()
}

func (m *mouseState) IsReleased(button mouse.Button) bool {
	return m.Button == button && m.Direction == mouse.DirRelease
}

func (m *mouseState) IsPressed(button mouse.Button) bool {
	return m.Button == button && m.Direction == mouse.DirPress
}

func (evt *mouseState) GetCursorPoint() CursorPoint {
	return &cursorPoint{int32(evt.X), int32(evt.Y)}
}

type CursorPoint interface {
	GetX() int32
	GetY() int32
	To() (int32, int32)
}
type cursorPoint struct {
	x, y int32
}

func (p *cursorPoint) GetX() int32 {
	return p.x
}
func (p *cursorPoint) GetY() int32 {
	return p.y
}

func (p *cursorPoint) To() (int32, int32) {
	return p.x, p.y
}
