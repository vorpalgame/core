package bus

import "github.com/vorpalgame/core/lib"

func NewPlayAudioEvent(state lib.AudioState) PlayAudioEvent {

	return &playAudioEventData{state.GetAudioFile(), state.IsAudioOnLoop()}
}

func NewStopAudioEvent(state lib.AudioState) StopAudioEvent {
	return &stopAudioEventData{state.GetAudioFile()}
}

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}

//////Basic AudioEvent //////

type AudioEvent interface {
	GetAudioFile() string
}

/////PlayAudioEvent
//Need to ensure there is asymmetery in events/implementations they can be
//distinguished by the case switch.

type PlayAudioEvent interface {
	AudioEvent
	IsLoop() bool
}

type playAudioEventData struct {
	fileName string
	isLoop   bool
}

func (p *playAudioEventData) GetAudioFile() string {
	return p.fileName
}

func (p *playAudioEventData) IsLoop() bool {
	return p.isLoop
}

/////StopAudioEvent

type StopAudioEvent interface {
	AudioEvent
}

type stopAudioEventData struct {
	fileName string
}

func (s *stopAudioEventData) GetAudioFile() string {
	return s.fileName
}
