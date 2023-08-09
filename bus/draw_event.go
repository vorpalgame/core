package bus

import "github.com/vorpalgame/core/lib"

///// Constructors //////////////////

func NewDrawLayersEvent() DrawLayersEvent {
	evt := drawLayerEvent{}
	evt.imageLayers = make([]lib.ImageLayerData, 0)
	return &evt
}
func NewDrawEvent() DrawEvent {
	return &drawEvent{}
}

/////////////////////////////////////

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

// /////////// DrawEvent and the drawEvent struct are no-ops or signaling at best.
type DrawEvent interface {
}
type drawEvent struct{}

/////////////////////////////////////////////////

type DrawLayersEvent interface {
	DrawEvent
	Reset()
	GetImageLayers() []lib.ImageLayerData
	AddImageLayer(imgLayer lib.ImageLayerData) DrawEvent
}

type drawLayerEvent struct {
	imageLayers []lib.ImageLayerData
}

func (evt *drawLayerEvent) Reset() {
	evt.imageLayers = make([]lib.ImageLayerData, 0)
}

func (evt *drawLayerEvent) AddImageLayer(img lib.ImageLayerData) DrawEvent {
	evt.imageLayers = append(evt.imageLayers, img)

	return evt
}

func (evt *drawLayerEvent) GetImageLayers() []lib.ImageLayerData {
	return evt.imageLayers
}
