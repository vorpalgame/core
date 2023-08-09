package audio

import (
	"sync"
)

// Minimal synchronization limited to cache access.
var (
	audioMapMutex = sync.RWMutex{}
)

type AudioCache interface {
	LoadPlayer(file string)
	GetPlayer(file string) AudioPlayer
}
type audioCacheData struct {
	audioCache map[string]AudioPlayer
}

func NewAudioCache() AudioCache {
	cache := audioCacheData{}
	cache.audioCache = make(map[string]AudioPlayer, 100)
	return &cache
}

func (c *audioCacheData) GetPlayer(fileName string) AudioPlayer {
	var player AudioPlayer
	audioMapMutex.Lock()
	player = c.audioCache[fileName]
	audioMapMutex.Unlock()
	return player
}
func (c *audioCacheData) LoadPlayer(fileName string) {
	audioMapMutex.Lock()
	if c.audioCache[fileName] == nil {
		c.audioCache[fileName] = NewAudioPlayer().Load(fileName)
	}
	audioMapMutex.Unlock()
}
