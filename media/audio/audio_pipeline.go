package audio

import (
	"github.com/vorpalgame/core/bus"
	"log"
)

func (data *audioPipelineData) OnAudioEvent(inputChannel <-chan bus.AudioEvent) {
	for evt := range inputChannel {
		//log.Println("Received audio event...")
		data.audioInputChannel <- evt
	}
}

type audioPipelineData struct {
	audioInputChannel chan bus.AudioEvent
}

func NewAudioPipeline(audioCache *AudioCache) {
	audioInputChannel := make(chan bus.AudioEvent, 100)
	data := audioPipelineData{audioInputChannel}
	bus.GetVorpalBus().AddAudioEventListener(&data)
	go loadCacheControlAudioPipeline(audioCache, audioInputChannel)

}

var loadCacheControlAudioPipeline = func(cache *AudioCache, inputChannel chan bus.AudioEvent) {
	stopAudioChannel := make(chan bus.StopAudioEvent, 1)
	loadAudioChannel := make(chan bus.PlayAudioEvent, 1)
	go loadAudio(cache, loadAudioChannel)
	go stopAudio(cache, stopAudioChannel)
	for event := range inputChannel {
		//log.Println("Received audio event: ", event)
		//log.Println("Received audio event type: ", reflect.TypeOf(event))
		_, ok := event.(bus.PlayAudioEvent)
		log.Println("Is PlayAudioEvent: ", ok)

		switch evt := event.(type) {
		case bus.PlayAudioEvent:
			log.Println("PlayOnce audio event...")
			loadAudioChannel <- evt
		case bus.StopAudioEvent:
			log.Println("Stop audio event...")
			stopAudioChannel <- evt
		default:
			log.Println("Unknown audio event: ", evt)
		}
	}
}

var loadAudio = func(cache *AudioCache, inputChannel chan bus.PlayAudioEvent) {
	playAudioChannel := make(chan bus.PlayAudioEvent, 1)
	go playAudio(cache, playAudioChannel)
	for evt := range inputChannel {
		log.Println("Load Audio: ", evt)

		for evt := range inputChannel {
			log.Println("Load: " + evt.GetAudioFile())
			(*cache).LoadPlayer(evt.GetAudioFile())
			playAudioChannel <- evt
		}
	}
}

var playAudio = func(cache *AudioCache, inputChannel chan bus.PlayAudioEvent) {
	for evt := range inputChannel {
		log.Println("PlayOnce audio function:", evt)
		var player AudioPlayer
		for player == nil {
			player = (*cache).GetPlayer(evt.GetAudioFile())
			log.Println("Player ", player)
		}
		player.Stop()
		player.PlayOnce()
	}
}

var stopAudio = func(cache *AudioCache, inputChannel chan bus.StopAudioEvent) {
	for evt := range inputChannel {
		log.Println("Stop Audio: ", evt)
		player := (*cache).GetPlayer(evt.GetAudioFile())
		if player != nil {
			player.Stop()
		}
	}

}
