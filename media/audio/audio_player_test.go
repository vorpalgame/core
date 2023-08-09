package audio

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testFile = "../../samples/resources/zombiecide/attack.mp3"

func TestPlayOnce(t *testing.T) {

	player := NewAudioPlayer().Load(testFile)
	player.PlayOnce()
	//Wait for up to five seconds.
	i := 0
	for i < 5 && player.IsPlaying() {
		time.Sleep(1 * time.Duration(time.Second))
		i++
	}
	assert.True(t, player.IsStopped())
}

func TestPlayLooped(t *testing.T) {
	player := NewAudioPlayer().Load(testFile)
	player.PlayLooped()
	//Wait for up to five seconds.
	i := 0
	for i < 10 && player.IsPlaying() {
		time.Sleep(1 * time.Duration(time.Second))
		i++
	}
	assert.True(t, player.IsPlaying())
	player.Stop()
	assert.True(t, player.IsStopped())
}
